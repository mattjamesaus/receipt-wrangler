package services

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/repositories"
	"receipt-wrangler/api/internal/structs"
)

var (
	ErrSupplierProfileNotFound    = repositories.ErrSupplierProfileNotFound
	ErrSupplierNameCollision      = repositories.ErrSupplierNameCollision
	ErrSupplierCategoryTagDenied  = errors.New("one or more categories or tags are not permitted")
	ErrSupplierCategoryTagMissing = errors.New("one or more categories or tags do not exist")
)

type SupplierProfileService struct {
	BaseService
}

func NewSupplierProfileService(tx *gorm.DB) SupplierProfileService {
	return SupplierProfileService{BaseService: BaseService{
		DB: repositories.GetDB(),
		TX: tx,
	}}
}

func (service SupplierProfileService) List(userId uint, groupId uint) ([]models.SupplierProfile, error) {
	profiles, err := repositories.NewSupplierProfileRepository(service.TX).GetByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	return service.filterProfilesForUser(userId, groupId, profiles)
}

func (service SupplierProfileService) Get(userId uint, groupId uint, profileId uint) (models.SupplierProfile, error) {
	profile, err := repositories.NewSupplierProfileRepository(service.TX).GetByIdInGroup(profileId, groupId)
	if err != nil {
		return models.SupplierProfile{}, err
	}
	filtered, err := service.filterProfilesForUser(userId, groupId, []models.SupplierProfile{profile})
	if err != nil {
		return models.SupplierProfile{}, err
	}
	return filtered[0], nil
}

func (service SupplierProfileService) Resolve(userId uint, groupId uint, name string) (*models.SupplierProfile, error) {
	normalised := NormaliseSupplierName(name)
	if len(normalised) == 0 {
		return nil, nil
	}

	profiles, err := repositories.NewSupplierProfileRepository(service.TX).FindEnabledByNormalisedName(groupId, normalised)
	if err != nil {
		return nil, err
	}
	if len(profiles) != 1 {
		return nil, nil
	}

	filtered, err := service.filterProfilesForUser(userId, groupId, profiles)
	if err != nil {
		return nil, err
	}
	return &filtered[0], nil
}

func (service SupplierProfileService) Create(userId uint, groupId uint, command commands.UpsertSupplierProfileCommand) (models.SupplierProfile, structs.ValidatorError, error) {
	profile, vErr, err := service.buildProfile(userId, groupId, 0, command)
	if err != nil || len(vErr.Errors) > 0 {
		return models.SupplierProfile{}, vErr, err
	}

	createdBy := userId
	profile.CreatedBy = &createdBy

	created, err := repositories.NewSupplierProfileRepository(service.TX).Create(profile)
	if err != nil {
		vErr, err := collisionOrError(err)
		return models.SupplierProfile{}, vErr, err
	}

	filtered, err := service.filterProfilesForUser(userId, groupId, []models.SupplierProfile{created})
	if err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}
	return filtered[0], structs.ValidatorError{}, nil
}

func (service SupplierProfileService) Update(userId uint, groupId uint, profileId uint, command commands.UpsertSupplierProfileCommand) (models.SupplierProfile, structs.ValidatorError, error) {
	existing, err := repositories.NewSupplierProfileRepository(service.TX).GetByIdInGroup(profileId, groupId)
	if err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}

	profile, vErr, err := service.buildProfile(userId, groupId, profileId, command)
	if err != nil || len(vErr.Errors) > 0 {
		return models.SupplierProfile{}, vErr, err
	}
	profile.ID = existing.ID
	profile.CreatedBy = existing.CreatedBy
	profile.CreatedByString = existing.CreatedByString
	profile.CreatedAt = existing.CreatedAt

	if err := service.preserveHiddenCatalog(userId, groupId, existing, &profile); err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}

	updated, err := repositories.NewSupplierProfileRepository(service.TX).Update(profile)
	if err != nil {
		vErr, err := collisionOrError(err)
		return models.SupplierProfile{}, vErr, err
	}

	filtered, err := service.filterProfilesForUser(userId, groupId, []models.SupplierProfile{updated})
	if err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}
	return filtered[0], structs.ValidatorError{}, nil
}

func (service SupplierProfileService) Delete(groupId uint, profileId uint) error {
	return repositories.NewSupplierProfileRepository(service.TX).Delete(profileId, groupId)
}

// ApplyAutoDefaults merges an auto-apply supplier profile into a receipt
// command about to be created. Categories and tags are added without removing
// existing selections. Expected currency is applied only when the command has
// no document currency — an extracted or caller-supplied value wins.
func (service SupplierProfileService) ApplyAutoDefaults(userId uint, command *commands.UpsertReceiptCommand) error {
	if command == nil || command.GroupId == 0 || len(strings.TrimSpace(command.Name)) == 0 {
		return nil
	}

	profile, err := service.Resolve(userId, command.GroupId, command.Name)
	if err != nil {
		return err
	}
	if profile == nil || !profile.AutoApply {
		return nil
	}

	command.Categories = mergeCategoryCommands(command.Categories, profile.Categories)
	command.Tags = mergeTagCommands(command.Tags, profile.Tags)

	if len(strings.TrimSpace(command.DocumentCurrencyCode)) == 0 &&
		profile.ExpectedDocumentCurrencyCode != nil &&
		len(strings.TrimSpace(*profile.ExpectedDocumentCurrencyCode)) > 0 {
		command.DocumentCurrencyCode = strings.ToUpper(strings.TrimSpace(*profile.ExpectedDocumentCurrencyCode))
	}

	return nil
}

func (service SupplierProfileService) buildProfile(userId uint, groupId uint, excludeProfileId uint, command commands.UpsertSupplierProfileCommand) (models.SupplierProfile, structs.ValidatorError, error) {
	vErr := structs.ValidatorError{Errors: map[string]string{}}

	displayName := strings.TrimSpace(command.Name)
	normalisedName := NormaliseSupplierName(displayName)
	if len(normalisedName) == 0 {
		vErr.Errors["name"] = "Name is empty after normalisation"
		return models.SupplierProfile{}, vErr, nil
	}

	aliases := uniqueNormalisedAliases(command.Aliases)
	for _, alias := range aliases {
		if alias.NormalisedName == normalisedName {
			vErr.Errors["aliases"] = "An alias cannot match the supplier name"
			return models.SupplierProfile{}, vErr, nil
		}
	}

	namesToCheck := make([]string, 0, 1+len(aliases))
	namesToCheck = append(namesToCheck, normalisedName)
	for _, alias := range aliases {
		namesToCheck = append(namesToCheck, alias.NormalisedName)
	}

	collisions, err := repositories.NewSupplierProfileRepository(service.TX).FindCollidingNormalisedNames(groupId, namesToCheck, excludeProfileId)
	if err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}
	if len(collisions) > 0 {
		vErr.Errors["name"] = fmt.Sprintf("A supplier profile or alias already uses the name %q in this group", collisions[0])
		return models.SupplierProfile{}, vErr, nil
	}

	categories, tags, err := service.loadAndAuthorizeCatalog(userId, groupId, command.CategoryIds, command.TagIds)
	if err != nil {
		return models.SupplierProfile{}, structs.ValidatorError{}, err
	}

	enabled := true
	if command.Enabled != nil {
		enabled = *command.Enabled
	}
	autoApply := false
	if command.AutoApply != nil {
		autoApply = *command.AutoApply
	}

	profileAliases := make([]models.SupplierProfileAlias, 0, len(aliases))
	for _, alias := range aliases {
		profileAliases = append(profileAliases, models.SupplierProfileAlias{
			GroupId:        groupId,
			Name:           alias.Name,
			NormalisedName: alias.NormalisedName,
		})
	}

	return models.SupplierProfile{
		GroupId:                      groupId,
		Name:                         displayName,
		NormalisedName:               normalisedName,
		ExpectedDocumentCurrencyCode: command.ExpectedDocumentCurrencyCode,
		Enabled:                      enabled,
		AutoApply:                    autoApply,
		Categories:                   categories,
		Tags:                         tags,
		Aliases:                      profileAliases,
	}, structs.ValidatorError{Errors: map[string]string{}}, nil
}

func (service SupplierProfileService) loadAndAuthorizeCatalog(userId uint, groupId uint, categoryIds []uint, tagIds []uint) ([]models.Category, []models.Tag, error) {
	categoryRepo := repositories.NewCategoryRepository(service.TX)
	tagRepo := repositories.NewTagsRepository(service.TX)

	uniqueCategoryIds := uniqueUintIds(categoryIds)
	uniqueTagIds := uniqueUintIds(tagIds)

	if len(uniqueCategoryIds) > 0 {
		count, err := categoryRepo.CountByIds(uniqueCategoryIds)
		if err != nil {
			return nil, nil, err
		}
		if count != int64(len(uniqueCategoryIds)) {
			return nil, nil, ErrSupplierCategoryTagMissing
		}
	}
	if len(uniqueTagIds) > 0 {
		count, err := tagRepo.CountByIds(uniqueTagIds)
		if err != nil {
			return nil, nil, err
		}
		if count != int64(len(uniqueTagIds)) {
			return nil, nil, ErrSupplierCategoryTagMissing
		}
	}

	permissionService := NewPermissionService(service.TX)
	allowed, err := permissionService.ValidateCategoryTagSelection(userId, groupId, uniqueCategoryIds, uniqueTagIds)
	if err != nil {
		return nil, nil, err
	}
	if !allowed {
		return nil, nil, ErrSupplierCategoryTagDenied
	}

	categories, err := categoryRepo.GetByIds(uniqueCategoryIds)
	if err != nil {
		return nil, nil, err
	}
	tags, err := tagRepo.GetByIds(uniqueTagIds)
	if err != nil {
		return nil, nil, err
	}

	return categories, tags, nil
}

func (service SupplierProfileService) filterProfilesForUser(userId uint, groupId uint, profiles []models.SupplierProfile) ([]models.SupplierProfile, error) {
	permissionService := NewPermissionService(service.TX)

	for i := range profiles {
		visibleCategories, err := permissionService.GetVisibleCategoriesForUser(userId, groupId, profiles[i].Categories)
		if err != nil {
			return nil, err
		}
		visibleTags, err := permissionService.GetVisibleTagsForUser(userId, groupId, profiles[i].Tags)
		if err != nil {
			return nil, err
		}
		profiles[i].Categories = visibleCategories
		profiles[i].Tags = visibleTags
	}

	return profiles, nil
}

// preserveHiddenCatalog keeps category/tag defaults the caller cannot see so a
// restricted editor cannot strip them by saving only the visible subset.
func (service SupplierProfileService) preserveHiddenCatalog(userId uint, groupId uint, existing models.SupplierProfile, profile *models.SupplierProfile) error {
	permissionService := NewPermissionService(service.TX)

	visibleCategories, err := permissionService.GetVisibleCategoriesForUser(userId, groupId, existing.Categories)
	if err != nil {
		return err
	}
	visibleTags, err := permissionService.GetVisibleTagsForUser(userId, groupId, existing.Tags)
	if err != nil {
		return err
	}

	visibleCategoryIds := make(map[uint]struct{}, len(visibleCategories))
	for _, category := range visibleCategories {
		visibleCategoryIds[category.ID] = struct{}{}
	}
	for _, category := range existing.Categories {
		if _, visible := visibleCategoryIds[category.ID]; !visible {
			profile.Categories = append(profile.Categories, category)
		}
	}

	visibleTagIds := make(map[uint]struct{}, len(visibleTags))
	for _, tag := range visibleTags {
		visibleTagIds[tag.ID] = struct{}{}
	}
	for _, tag := range existing.Tags {
		if _, visible := visibleTagIds[tag.ID]; !visible {
			profile.Tags = append(profile.Tags, tag)
		}
	}

	return nil
}

func collisionOrError(err error) (structs.ValidatorError, error) {
	if errors.Is(err, ErrSupplierNameCollision) {
		return structs.ValidatorError{Errors: map[string]string{
			"name": "A supplier profile or alias already uses this name in this group",
		}}, nil
	}
	return structs.ValidatorError{}, err
}

func mergeCategoryCommands(existing []commands.UpsertCategoryCommand, additions []models.Category) []commands.UpsertCategoryCommand {
	seen := make(map[uint]struct{}, len(existing))
	for _, category := range existing {
		if category.Id != nil {
			seen[*category.Id] = struct{}{}
		}
	}
	result := append([]commands.UpsertCategoryCommand{}, existing...)
	for _, category := range additions {
		if _, exists := seen[category.ID]; exists {
			continue
		}
		id := category.ID
		result = append(result, commands.UpsertCategoryCommand{
			Id:          &id,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return result
}

func mergeTagCommands(existing []commands.UpsertTagCommand, additions []models.Tag) []commands.UpsertTagCommand {
	seen := make(map[uint]struct{}, len(existing))
	for _, tag := range existing {
		if tag.Id != nil {
			seen[*tag.Id] = struct{}{}
		}
	}
	result := append([]commands.UpsertTagCommand{}, existing...)
	for _, tag := range additions {
		if _, exists := seen[tag.ID]; exists {
			continue
		}
		id := tag.ID
		result = append(result, commands.UpsertTagCommand{
			Id:          &id,
			Name:        tag.Name,
			Description: tag.Description,
		})
	}
	return result
}

func uniqueUintIds(ids []uint) []uint {
	seen := make(map[uint]struct{}, len(ids))
	result := make([]uint, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
		result = append(result, id)
	}
	return result
}
