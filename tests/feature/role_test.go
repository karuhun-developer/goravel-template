package feature

import (
	"bytes"
	"encoding/json"
	"testing"
	"fmt"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"
	"karuhundeveloper.com/gostarterkit/app/models/user"
	"karuhundeveloper.com/gostarterkit/app/models/role"
	"karuhundeveloper.com/gostarterkit/tests"
)

type RoleTestSuite struct {
	suite.Suite
	tests.TestCase
	user  *user.User
	token string
}

func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}

func (s *RoleTestSuite) SetupTest() {
	password, err := facades.Hash().Make("password")
	s.Nil(err)

	s.user = &user.User{
		Name:     "Test User Role",
		Email:    "test_role@example.com",
		Password: password,
	}
	s.Nil(facades.Orm().Query().Create(s.user))

	// Create Permissions
	permissions := []string{
		"create_role", "view_role", "update_role", "delete_role",
		"create_permission", "view_permission", "update_permission", "delete_permission",
		"view_all_role",
	}
	var createdPermissions []*role.Permission
	for _, p := range permissions {
		perm := &role.Permission{Name: p}
		s.Nil(facades.Orm().Query().Create(perm))
		createdPermissions = append(createdPermissions, perm)
	}

	// Create Role
	adminRole := &role.Role{Name: "Admin"}
	s.Nil(facades.Orm().Query().Create(adminRole))
	s.Nil(facades.Orm().Query().Where("name", "Admin").First(adminRole))
	fmt.Printf("Admin Role ID: %d\n", adminRole.ID)

	// Assign Permissions to Role
	for _, perm := range createdPermissions {
		s.Nil(facades.Orm().Query().Where("name", perm.Name).First(perm))
		fmt.Printf("Assigning Permission ID: %d to Role ID: %d\n", perm.ID, adminRole.ID)
		rolePerm := &role.RolePermission{
			RoleID:       adminRole.ID,
			PermissionID: perm.ID,
		}
		s.Nil(facades.Orm().Query().Create(rolePerm))
	}

	// Assign Role to User
	userRole := &user.UserRole{
		UserId: s.user.ID,
		RoleId: adminRole.ID,
	}
	s.Nil(facades.Orm().Query().Create(userRole))
	
	var userRoles []user.UserRole
	facades.Orm().Query().Where("user_id", s.user.ID).Where("role_id", adminRole.ID).Get(&userRoles)
	fmt.Printf("User Role Count: %d for User ID: %d Role ID: %d\n", len(userRoles), s.user.ID, adminRole.ID)

	// Manual Permission Check
	var rp role.RolePermission
	err = facades.Orm().Query().Model(&role.RolePermission{}).
		Join("join roles on roles.id = role_permissions.role_id").
		Join("join permissions on permissions.id = role_permissions.permission_id").
		Where("roles.id", adminRole.ID).
		Where("permissions.name", "create_role").
		Select("role_permissions.*").
		First(&rp)
	if err != nil {
		fmt.Printf("HasPermission check failed: %v\n", err)
	} else {
		fmt.Printf("HasPermission check passed: %+v\n", rp)
	}

	// Login
	payload := map[string]any{
		"email":    "test_role@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)
	response, err := s.Http(s.T()).Post("/api/v1/auth/login", bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(200)

	data, err := response.Json()
	s.Nil(err)
	s.token = data["data"].(map[string]any)["token"].(string)
}

func (s *RoleTestSuite) TearDownTest() {
	facades.Orm().Query().Delete(s.user)
	facades.Orm().Query().Where("name", "Admin").Delete(&role.Role{})
	facades.Orm().Query().Where("name", "Test Role").Delete(&role.Role{})
	facades.Orm().Query().Where("name", "Updated Role").Delete(&role.Role{})
	
	permissions := []string{
		"create_role", "view_role", "update_role", "delete_role",
		"create_permission", "view_permission", "update_permission", "delete_permission",
		"Test Permission", "Updated Permission", "view_all_role",
	}
	facades.Orm().Query().Where("name IN ?", permissions).Delete(&role.Permission{})
}

func (s *RoleTestSuite) TestRoleCRUD() {
	// Create
	payload := map[string]any{
		"name": "Test Role",
	}
	body, _ := json.Marshal(payload)
	response, err := s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Post("/api/v1/role/roles", bytes.NewBuffer(body))
	s.Nil(err)
	content, _ := response.Content()
	fmt.Printf("Create Role Response: %s\n", content)
	response.AssertStatus(201)
	
	data, err := response.Json()
	s.Nil(err)
	roleID := data["data"].(map[string]any)["id"].(float64)

	// Read Index
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Get("/api/v1/role/roles")
	s.Nil(err)
	response.AssertStatus(200)

	// Read Show
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Get("/api/v1/role/roles/" + s.getIDString(roleID))
	s.Nil(err)
	response.AssertStatus(200)

	// Update
	payload = map[string]any{
		"name": "Updated Role",
	}
	body, _ = json.Marshal(payload)
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Put("/api/v1/role/roles/"+s.getIDString(roleID), bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(200)

	// Delete
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Delete("/api/v1/role/roles/"+s.getIDString(roleID), nil)
	s.Nil(err)
	response.AssertStatus(200)
}

func (s *RoleTestSuite) TestPermissionCRUD() {
	// Create
	payload := map[string]any{
		"name": "Test Permission",
	}
	body, _ := json.Marshal(payload)
	response, err := s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Post("/api/v1/role/permissions", bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(201)

	data, err := response.Json()
	s.Nil(err)
	permissionID := data["data"].(map[string]any)["id"].(float64)

	// Read Index
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Get("/api/v1/role/permissions")
	s.Nil(err)
	response.AssertStatus(200)

	// Read Show
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Get("/api/v1/role/permissions/" + s.getIDString(permissionID))
	s.Nil(err)
	response.AssertStatus(200)

	// Update
	payload = map[string]any{
		"name": "Updated Permission",
	}
	body, _ = json.Marshal(payload)
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Put("/api/v1/role/permissions/"+s.getIDString(permissionID), bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(200)

	// Delete
	response, err = s.Http(s.T()).WithHeader("Authorization", "Bearer "+s.token).Delete("/api/v1/role/permissions/"+s.getIDString(permissionID), nil)
	s.Nil(err)
	response.AssertStatus(200)
}

func (s *RoleTestSuite) getIDString(id float64) string {
	return fmt.Sprintf("%.0f", id)
}
