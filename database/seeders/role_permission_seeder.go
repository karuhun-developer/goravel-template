package seeders

import (
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/models/role"
)

type RolePermissionSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *RolePermissionSeeder) Signature() string {
	return "RolePermissionSeeder"
}

// List prefix of permissions
var permissionList = [...]string{
	"view_all_",
	"view_",
	"create_",
	"edit_",
	"delete_",
	"export_",
	"import_",
	"approval_",
	// "force_delete_",
	// "restore_",
}

// Permission list of resources
var resources = [...]string{
	"user",
	"role",
	"permission",
	"role_permission",
}

// User permissions
var userPermissions = [...]string{
	"view_user",
}


// Run executes the seeder logic.
func (s *RolePermissionSeeder) Run() error {
	var superadminRole role.Role
	var userRole role.Role

	// First or create roles
	facades.Orm().Query().Where("name", "superadmin").FirstOrCreate(&superadminRole, role.Role{Name: "superadmin"})
	facades.Orm().Query().Where("name", "user").FirstOrCreate(&userRole, role.Role{Name: "user"})

	// Clear cache
	facades.Cache().Flush()

	// Clear the associations
	facades.Orm().Query().Model(&superadminRole).Association("Permissions").Clear()
	facades.Orm().Query().Model(&userRole).Association("Permissions").Clear()

	// Define permissions
	for _, resource := range resources {
		for _, permPrefix := range permissionList {
			permissionName := permPrefix + resource

			var permission role.Permission
			facades.Orm().Query().Where("name", permissionName).FirstOrCreate(&permission, role.Permission{Name: permissionName})

			// Assign all permissions to superadmin
			facades.Orm().Query().Model(&superadminRole).Association("Permissions").Append(&permission)

			// Assign specific permissions to user role
			for _, userPerm := range userPermissions {
				if permissionName == userPerm {
					facades.Orm().Query().Model(&userRole).Association("Permissions").Append(&permission)
				}
			}
		}
	}


	return nil
}
