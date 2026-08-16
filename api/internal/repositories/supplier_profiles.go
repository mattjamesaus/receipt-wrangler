package repositories

import (
	"errors"

	"gorm.io/gorm"
	"receipt-wrangler/api/internal/models"
)

var (
	ErrSupplierProfileNotFound = errors.New("supplier profile not found")
	ErrSupplierNameCollision   = errors.New("supplier name already in use")
)

type SupplierProfileRepository struct {
	BaseRepository
}

func NewSupplierProfileRepository(tx *gorm.DB) SupplierProfileRepository {
	return SupplierProfileRepository{BaseRepository: BaseRepository{
		DB: GetDB(),
		TX: tx,
	}}
}

func (repository SupplierProfileRepository) preload(db *gorm.DB) *gorm.DB {
	return db.Preload("Categories").Preload("Tags").Preload("Aliases")
}

func (repository SupplierProfileRepository) GetByGroupId(groupId uint) ([]models.SupplierProfile, error) {
	var profiles []models.SupplierProfile
	err := repository.preload(repository.GetDB()).
		Where("group_id = ?", groupId).
		Order("name ASC").
		Find(&profiles).Error
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (repository SupplierProfileRepository) GetByIdInGroup(profileId uint, groupId uint) (models.SupplierProfile, error) {
	var profile models.SupplierProfile
	err := repository.preload(repository.GetDB()).
		Where("id = ? AND group_id = ?", profileId, groupId).
		First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SupplierProfile{}, ErrSupplierProfileNotFound
		}
		return models.SupplierProfile{}, err
	}
	return profile, nil
}

func (repository SupplierProfileRepository) FindEnabledByNormalisedName(groupId uint, normalisedName string) ([]models.SupplierProfile, error) {
	if len(normalisedName) == 0 {
		return []models.SupplierProfile{}, nil
	}

	db := repository.GetDB()
	var profileIds []uint

	err := db.Model(&models.SupplierProfile{}).
		Select("id").
		Where("group_id = ? AND enabled = ? AND normalised_name = ?", groupId, true, normalisedName).
		Pluck("id", &profileIds).Error
	if err != nil {
		return nil, err
	}

	var aliasProfileIds []uint
	err = db.Model(&models.SupplierProfileAlias{}).
		Select("supplier_profile_id").
		Where("group_id = ? AND normalised_name = ?", groupId, normalisedName).
		Pluck("supplier_profile_id", &aliasProfileIds).Error
	if err != nil {
		return nil, err
	}

	idSet := make(map[uint]struct{})
	for _, id := range profileIds {
		idSet[id] = struct{}{}
	}
	for _, id := range aliasProfileIds {
		idSet[id] = struct{}{}
	}
	if len(idSet) == 0 {
		return []models.SupplierProfile{}, nil
	}

	ids := make([]uint, 0, len(idSet))
	for id := range idSet {
		ids = append(ids, id)
	}

	var profiles []models.SupplierProfile
	err = repository.preload(db).
		Where("id IN ? AND group_id = ? AND enabled = ?", ids, groupId, true).
		Find(&profiles).Error
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

// FindCollidingNormalisedNames returns any of the supplied normalised names
// that already belong to a profile or alias in the group, excluding excludeProfileId
// (0 means do not exclude).
func (repository SupplierProfileRepository) FindCollidingNormalisedNames(groupId uint, normalisedNames []string, excludeProfileId uint) ([]string, error) {
	return repository.findCollidingNormalisedNames(repository.GetDB(), groupId, normalisedNames, excludeProfileId)
}

func (repository SupplierProfileRepository) findCollidingNormalisedNames(db *gorm.DB, groupId uint, normalisedNames []string, excludeProfileId uint) ([]string, error) {
	if len(normalisedNames) == 0 {
		return nil, nil
	}

	collisions := make(map[string]struct{})

	var profileNames []string
	query := db.Model(&models.SupplierProfile{}).
		Select("normalised_name").
		Where("group_id = ? AND normalised_name IN ?", groupId, normalisedNames)
	if excludeProfileId > 0 {
		query = query.Where("id <> ?", excludeProfileId)
	}
	if err := query.Pluck("normalised_name", &profileNames).Error; err != nil {
		return nil, err
	}
	for _, name := range profileNames {
		collisions[name] = struct{}{}
	}

	var aliasNames []string
	aliasQuery := db.Model(&models.SupplierProfileAlias{}).
		Select("normalised_name").
		Where("group_id = ? AND normalised_name IN ?", groupId, normalisedNames)
	if excludeProfileId > 0 {
		aliasQuery = aliasQuery.Where("supplier_profile_id <> ?", excludeProfileId)
	}
	if err := aliasQuery.Pluck("normalised_name", &aliasNames).Error; err != nil {
		return nil, err
	}
	for _, name := range aliasNames {
		collisions[name] = struct{}{}
	}

	result := make([]string, 0, len(collisions))
	for name := range collisions {
		result = append(result, name)
	}
	return result, nil
}

func (repository SupplierProfileRepository) assertNoCollisions(tx *gorm.DB, profile models.SupplierProfile) error {
	names := make([]string, 0, 1+len(profile.Aliases))
	names = append(names, profile.NormalisedName)
	for _, alias := range profile.Aliases {
		names = append(names, alias.NormalisedName)
	}

	collisions, err := repository.findCollidingNormalisedNames(tx, profile.GroupId, names, profile.ID)
	if err != nil {
		return err
	}
	if len(collisions) > 0 {
		return ErrSupplierNameCollision
	}
	return nil
}

func (repository SupplierProfileRepository) Create(profile models.SupplierProfile) (models.SupplierProfile, error) {
	db := repository.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := repository.assertNoCollisions(tx, profile); err != nil {
			return err
		}

		aliases := profile.Aliases
		categories := profile.Categories
		tags := profile.Tags
		enabled := profile.Enabled
		profile.Aliases = nil
		profile.Categories = nil
		profile.Tags = nil
		if err := tx.Omit("Categories", "Tags", "Aliases").Create(&profile).Error; err != nil {
			return err
		}
		// GORM omits the false zero-value on Create, so a disabled profile
		// would otherwise keep the column default of true.
		if err := tx.Model(&models.SupplierProfile{}).
			Where("id = ?", profile.ID).
			Update("enabled", enabled).Error; err != nil {
			return err
		}
		profile.Enabled = enabled
		for i := range aliases {
			aliases[i].SupplierProfileId = profile.ID
			aliases[i].GroupId = profile.GroupId
		}
		if len(aliases) > 0 {
			if err := tx.Create(&aliases).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&profile).Association("Categories").Replace(categories); err != nil {
			return err
		}
		if err := tx.Model(&profile).Association("Tags").Replace(tags); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.SupplierProfile{}, err
	}

	return repository.GetByIdInGroup(profile.ID, profile.GroupId)
}

func (repository SupplierProfileRepository) Update(profile models.SupplierProfile) (models.SupplierProfile, error) {
	db := repository.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := repository.assertNoCollisions(tx, profile); err != nil {
			return err
		}

		if err := tx.Model(&models.SupplierProfile{}).Where("id = ? AND group_id = ?", profile.ID, profile.GroupId).Updates(map[string]interface{}{
			"name":                            profile.Name,
			"normalised_name":                 profile.NormalisedName,
			"expected_document_currency_code": profile.ExpectedDocumentCurrencyCode,
			"enabled":                         profile.Enabled,
		}).Error; err != nil {
			return err
		}

		if err := tx.Where("supplier_profile_id = ?", profile.ID).Delete(&models.SupplierProfileAlias{}).Error; err != nil {
			return err
		}
		for i := range profile.Aliases {
			profile.Aliases[i].SupplierProfileId = profile.ID
			profile.Aliases[i].GroupId = profile.GroupId
		}
		if len(profile.Aliases) > 0 {
			if err := tx.Create(&profile.Aliases).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&profile).Association("Categories").Replace(profile.Categories); err != nil {
			return err
		}
		if err := tx.Model(&profile).Association("Tags").Replace(profile.Tags); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.SupplierProfile{}, err
	}

	return repository.GetByIdInGroup(profile.ID, profile.GroupId)
}

func (repository SupplierProfileRepository) Delete(profileId uint, groupId uint) error {
	db := repository.GetDB()

	return db.Transaction(func(tx *gorm.DB) error {
		var profile models.SupplierProfile
		if err := tx.Where("id = ? AND group_id = ?", profileId, groupId).First(&profile).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrSupplierProfileNotFound
			}
			return err
		}

		if err := tx.Where("supplier_profile_id = ?", profileId).Delete(&models.SupplierProfileAlias{}).Error; err != nil {
			return err
		}
		if err := tx.Where("supplier_profile_id = ?", profileId).Delete(&models.SupplierProfileCategory{}).Error; err != nil {
			return err
		}
		if err := tx.Where("supplier_profile_id = ?", profileId).Delete(&models.SupplierProfileTag{}).Error; err != nil {
			return err
		}
		return tx.Delete(&profile).Error
	})
}
