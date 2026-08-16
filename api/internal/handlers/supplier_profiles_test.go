package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/go-chi/chi/v5"
	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/permissions"
	"receipt-wrangler/api/internal/repositories"
	"receipt-wrangler/api/internal/services"
	"receipt-wrangler/api/internal/utils"
)

func tearDownSupplierProfileHandlerTest() {
	repositories.TruncateTestDb()
}

func supplierHandlerRequest(method string, groupId string, profileId string, userId uint, body string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, "/api", strings.NewReader(body))
	route := chi.NewRouteContext()
	route.URLParams.Add("groupId", groupId)
	if len(profileId) > 0 {
		route.URLParams.Add("profileId", profileId)
	}
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, route))
	r = r.WithContext(context.WithValue(r.Context(), jwtmiddleware.ContextKey{}, claimsForUser(userId)))
	return w, r
}

func TestCreateAndResolveSupplierProfile(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)

	w, r := supplierHandlerRequest("POST", "1", "", 1, `{"name":"GitHub","aliases":["GitHub, Inc."],"categoryIds":[1],"expectedDocumentCurrencyCode":"USD"}`)
	CreateSupplierProfile(w, r)
	if w.Result().StatusCode != 200 {
		utils.PrintTestError(t, w.Body.String(), "200 create")
		return
	}

	var created models.SupplierProfile
	if err := json.Unmarshal(w.Body.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	if created.Name != "GitHub" || !created.Enabled {
		t.Fatalf("created = %#v", created)
	}

	w, r = supplierHandlerRequest("POST", "1", "", 1, `{"name":"GitHub, Inc."}`)
	ResolveSupplierProfile(w, r)
	if w.Result().StatusCode != 200 {
		utils.PrintTestError(t, w.Body.String(), "200 resolve")
		return
	}

	var resolved commands.ResolveSupplierProfileResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resolved); err != nil {
		t.Fatal(err)
	}
	if resolved.Profile == nil || resolved.Profile.ID != created.ID {
		t.Fatalf("resolve = %#v", resolved.Profile)
	}
}

func TestCreateSupplierProfileRequiresUpdatePermission(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate)

	w, r := supplierHandlerRequest("POST", "1", "", 1, `{"name":"GitHub","categoryIds":[1]}`)
	CreateSupplierProfile(w, r)
	if w.Result().StatusCode != 403 {
		utils.PrintTestError(t, w.Result().StatusCode, 403)
	}
}

func TestListSupplierProfilesRequiresCreatePermission(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsRead)

	w, r := supplierHandlerRequest("GET", "1", "", 1, "")
	ListSupplierProfiles(w, r)
	if w.Result().StatusCode != 403 {
		utils.PrintTestError(t, w.Result().StatusCode, 403)
	}
}

func TestSupplierProfileCannotBeReadFromAnotherGroup(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)
	grantGroupPerms(t, 4, 2, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)

	created, vErr, err := services.NewSupplierProfileService(nil).Create(1, 1, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{1},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	w, r := supplierHandlerRequest("GET", "2", fmt.Sprintf("%d", created.ID), 4, "")
	GetSupplierProfile(w, r)
	if w.Result().StatusCode != 404 {
		utils.PrintTestError(t, w.Result().StatusCode, 404)
	}
}

func TestSupplierProfileHidesRestrictedCategories(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)

	created, vErr, err := services.NewSupplierProfileService(nil).Create(1, 1, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{1, 2},
	})
	if err != nil || len(vErr.Errors) > 0 {
		t.Fatalf("create: %v %#v", err, vErr.Errors)
	}

	role, err := repositories.NewRoleRepository(nil).CreateGroupRole(
		"Restricted Supplier Viewer",
		"",
		[]string{permissions.GroupReceiptsCreate},
		[]uint{1},
		nil,
		nil,
		false, false,
	)
	if err != nil {
		t.Fatal(err)
	}
	restricted := models.User{Username: "restricted-supplier", Password: "password"}
	if err := repositories.GetDB().Create(&restricted).Error; err != nil {
		t.Fatal(err)
	}
	if err := repositories.GetDB().Create(&models.GroupMember{GroupID: 1, UserID: restricted.ID, GroupRoleID: &role.ID}).Error; err != nil {
		t.Fatal(err)
	}

	w, r := supplierHandlerRequest("GET", "1", fmt.Sprintf("%d", created.ID), restricted.ID, "")
	GetSupplierProfile(w, r)
	if w.Result().StatusCode != 200 {
		utils.PrintTestError(t, w.Body.String(), "200")
		return
	}

	var profile models.SupplierProfile
	if err := json.Unmarshal(w.Body.Bytes(), &profile); err != nil {
		t.Fatal(err)
	}
	if len(profile.Categories) != 1 || profile.Categories[0].ID != 1 {
		t.Fatalf("expected only the granted category, got %#v", profile.Categories)
	}
}

func TestCreateSupplierProfileRejectsCollision(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)

	w, r := supplierHandlerRequest("POST", "1", "", 1, `{"name":"GitHub","categoryIds":[1]}`)
	CreateSupplierProfile(w, r)
	if w.Result().StatusCode != 200 {
		t.Fatalf("first create: %s", w.Body.String())
	}

	w, r = supplierHandlerRequest("POST", "1", "", 1, `{"name":"Other","aliases":["github"],"categoryIds":[1]}`)
	CreateSupplierProfile(w, r)
	if w.Result().StatusCode != 400 {
		utils.PrintTestError(t, w.Result().StatusCode, 400)
	}
}

func TestDeleteSupplierProfile(t *testing.T) {
	defer tearDownSupplierProfileHandlerTest()
	repositories.CreateTestGroupWithUsers()
	repositories.CreateTestCategories()
	grantGroupPerms(t, 1, 1, permissions.GroupReceiptsCreate, permissions.GroupReceiptsUpdate)

	created, _, err := services.NewSupplierProfileService(nil).Create(1, 1, commands.UpsertSupplierProfileCommand{
		Name:        "GitHub",
		CategoryIds: []uint{1},
	})
	if err != nil {
		t.Fatal(err)
	}

	w, r := supplierHandlerRequest("DELETE", "1", fmt.Sprintf("%d", created.ID), 1, "")
	DeleteSupplierProfile(w, r)
	if w.Result().StatusCode != 200 {
		utils.PrintTestError(t, w.Result().StatusCode, 200)
	}

	w, r = supplierHandlerRequest("GET", "1", fmt.Sprintf("%d", created.ID), 1, "")
	GetSupplierProfile(w, r)
	if w.Result().StatusCode != 404 {
		utils.PrintTestError(t, w.Result().StatusCode, 404)
	}
}
