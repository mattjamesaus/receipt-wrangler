package services

import (
	"errors"
	"testing"

	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/permissions"
	"receipt-wrangler/api/internal/repositories"
)

func tearDownSupplierProfileServiceTest() {
	repositories.TruncateTestDb()
}

func seedSupplierGroup(t *testing.T) (uint, uint, models.Category, models.Tag) {
	t.Helper()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()

	db := repositories.GetDB()
	tag := models.Tag{Name: "Work"}
	if err := db.Create(&tag).Error; err != nil {
		t.Fatalf("create tag: %v", err)
	}

	category := models.Category{}
	if err := db.First(&category, 1).Error; err != nil {
		t.Fatalf("load category: %v", err)
	}

	userId := uint(1)
	groupId := uint(1)
	grantSupplierGroupPerms(t, userId, groupId, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)
	return userId, groupId, category, tag
}

func grantSupplierGroupPerms(t *testing.T, userId uint, groupId uint, perms ...string) {
	t.Helper()
	ClearRolePermissionCacheForTests()
	ClearGroupRoleGrantCacheForTests()

	role, err := repositories.NewRoleRepository(nil).CreateGroupRole(
		"supplier-test-role",
		"",
		perms,
		nil, nil, nil,
		false, false,
	)
	if err != nil {
		t.Fatalf("create role: %v", err)
	}
	if err := repositories.GetDB().Model(&models.GroupMember{}).
		Where("user_id = ? AND group_id = ?", userId, groupId).
		Update("group_role_id", role.ID).Error; err != nil {
		t.Fatalf("assign role: %v", err)
	}
}

func TestSupplierProfileService_CreateAndResolveAlias(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, tag := seedSupplierGroup(t)
	currency := "USD"

	service := NewSupplierProfileService(nil)
	profile, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:                         "GitHub",
		Aliases:                      []string{"GitHub, Inc."},
		CategoryIds:                  []uint{category.ID},
		TagIds:                       []uint{tag.ID},
		ExpectedDocumentCurrencyCode: &currency,
	})
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if len(vErr.Errors) > 0 {
		t.Fatalf("create validation: %#v", vErr.Errors)
	}
	if profile.NormalisedName != "github" {
		t.Fatalf("normalised name = %q", profile.NormalisedName)
	}
	if len(profile.Aliases) != 1 || profile.Aliases[0].NormalisedName != "github inc" {
		t.Fatalf("aliases = %#v", profile.Aliases)
	}

	matched, err := service.Resolve(userId, groupId, "GitHub, Inc.")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if matched == nil || matched.ID != profile.ID {
		t.Fatalf("expected alias match, got %#v", matched)
	}

	matched, err = service.Resolve(userId, groupId, "  github,   inc. ")
	if err != nil {
		t.Fatalf("resolve variant: %v", err)
	}
	if matched == nil || matched.ID != profile.ID {
		t.Fatalf("expected punctuation/case match, got %#v", matched)
	}
}

func TestSupplierProfileService_RejectsCollision(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, _ := seedSupplierGroup(t)
	service := NewSupplierProfileService(nil)

	_, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		Aliases:     []string{"GitHub, Inc."},
		CategoryIds: []uint{category.ID},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("seed create: %v %#v", err, vErr.Errors)
	}

	_, vErr, err = service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "Other",
		Aliases:     []string{"github"},
		CategoryIds: []uint{category.ID},
	})
	if err != nil {
		t.Fatalf("collision create err: %v", err)
	}
	if vErr.Errors["name"] == "" {
		t.Fatalf("expected collision on alias vs canonical, got %#v", vErr.Errors)
	}
}

func TestSupplierProfileService_DisabledDoesNotMatch(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, _ := seedSupplierGroup(t)
	service := NewSupplierProfileService(nil)
	enabled := false

	profile, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID},
		Enabled:     &enabled,
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	matched, err := service.Resolve(userId, groupId, "GitHub")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if matched != nil {
		t.Fatalf("disabled profile should not match, got %#v", matched)
	}

	enabled = true
	_, vErr, err = service.Update(userId, groupId, profile.ID, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID},
		Enabled:     &enabled,
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("enable: %v %#v", err, vErr.Errors)
	}

	matched, err = service.Resolve(userId, groupId, "GitHub")
	if err != nil {
		t.Fatalf("resolve after enable: %v", err)
	}
	if matched == nil {
		t.Fatal("expected match after re-enable")
	}
}

func TestSupplierProfileService_GroupIsolation(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, _ := seedSupplierGroup(t)
	service := NewSupplierProfileService(nil)

	_, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	matched, err := service.Resolve(userId, 2, "GitHub")
	if err != nil {
		t.Fatalf("resolve other group: %v", err)
	}
	if matched != nil {
		t.Fatal("profile must not leak across groups")
	}

	_, err = service.Get(userId, 2, 1)
	if !errors.Is(err, ErrSupplierProfileNotFound) {
		t.Fatalf("expected not found across groups, got %v", err)
	}
}

func TestSupplierProfileService_UpdatePreservesHiddenCatalog(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, tag := seedSupplierGroup(t)
	service := NewSupplierProfileService(nil)

	second := models.Category{}
	if err := repositories.GetDB().First(&second, 2).Error; err != nil {
		t.Fatalf("load second category: %v", err)
	}

	profile, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID, second.ID},
		TagIds:      []uint{tag.ID},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	role, err := repositories.NewRoleRepository(nil).CreateGroupRole(
		"restricted-supplier-editor",
		"",
		[]string{permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate},
		[]uint{category.ID},
		nil,
		nil,
		false, false,
	)
	if err != nil {
		t.Fatalf("create restricted role: %v", err)
	}
	restricted := models.User{Username: "restricted-supplier-editor", Password: "password"}
	if err := repositories.GetDB().Create(&restricted).Error; err != nil {
		t.Fatalf("create restricted user: %v", err)
	}
	if err := repositories.GetDB().Create(&models.GroupMember{
		GroupID:     groupId,
		UserID:      restricted.ID,
		GroupRoleID: &role.ID,
	}).Error; err != nil {
		t.Fatalf("create membership: %v", err)
	}
	ClearRolePermissionCacheForTests()
	ClearGroupRoleGrantCacheForTests()

	got, err := service.Get(restricted.ID, groupId, profile.ID)
	if err != nil {
		t.Fatalf("get as restricted: %v", err)
	}
	if len(got.Categories) != 1 || got.Categories[0].ID != category.ID {
		t.Fatalf("restricted get should hide category 2, got %#v", got.Categories)
	}

	enabled := false
	updated, vErr, err := service.Update(restricted.ID, groupId, profile.ID, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID},
		TagIds:      []uint{tag.ID},
		Enabled:     &enabled,
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("restricted update: %v %#v", err, vErr.Errors)
	}
	if len(updated.Categories) != 1 || updated.Categories[0].ID != category.ID {
		t.Fatalf("restricted response should still hide category 2, got %#v", updated.Categories)
	}

	reloaded, err := service.Get(userId, groupId, profile.ID)
	if err != nil {
		t.Fatalf("reload as owner: %v", err)
	}
	if !containsCategoryID(reloaded.Categories, category.ID) || !containsCategoryID(reloaded.Categories, second.ID) {
		t.Fatalf("hidden category was stripped: %#v", reloaded.Categories)
	}
	if reloaded.Enabled {
		t.Fatal("expected profile to be disabled")
	}
}

func containsCategoryID(categories []models.Category, id uint) bool {
	for _, category := range categories {
		if category.ID == id {
			return true
		}
	}
	return false
}

func TestSupplierProfileService_CategoryDeleteUnlinks(t *testing.T) {
	defer tearDownSupplierProfileServiceTest()
	userId, groupId, category, tag := seedSupplierGroup(t)
	service := NewSupplierProfileService(nil)

	profile, vErr, err := service.Create(userId, groupId, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{category.ID},
		TagIds:      []uint{tag.ID},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	if err := repositories.NewCategoryRepository(nil).DeleteCategory(category.ID); err != nil {
		t.Fatalf("delete category: %v", err)
	}

	reloaded, err := service.Get(userId, groupId, profile.ID)
	if err != nil {
		t.Fatalf("get after category delete: %v", err)
	}
	if len(reloaded.Categories) != 0 {
		t.Fatalf("expected category unlinked, got %#v", reloaded.Categories)
	}
	if len(reloaded.Tags) != 1 {
		t.Fatalf("expected tag to remain, got %#v", reloaded.Tags)
	}
}
