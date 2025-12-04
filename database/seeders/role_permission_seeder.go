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
	"update_",
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
	"anggota",
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
	facades.Orm().Query().Model(&role.RolePermission{}).Where("role_id", superadminRole.ID).Delete(&role.RolePermission{})
	facades.Orm().Query().Model(&role.RolePermission{}).Where("role_id", userRole.ID).Delete(&role.RolePermission{})

	// Define permissions
	for _, resource := range resources {
		for _, permPrefix := range permissionList {
			permissionName := permPrefix + resource

			var permission role.Permission
			facades.Orm().Query().Where("name", permissionName).FirstOrCreate(&permission, role.Permission{Name: permissionName})

			// Assign all permissions to superadmin if not already associated
			exists, _ := facades.Orm().Query().Model(&role.RolePermission{}).
				Where("role_id", superadminRole.ID).
				Where("permission_id", permission.ID).
				Exists()
			if !exists {
				permissionRole := role.RolePermission{
					RoleID:       superadminRole.ID,
					PermissionID: permission.ID,
				}
				facades.Orm().Query().Create(&permissionRole)
			}

			// Assign specific permissions to user role
			for _, userPerm := range userPermissions {
				if permissionName == userPerm {
					// Check if association already exists
					exists, _ := facades.Orm().Query().Model(&role.RolePermission{}).
						Where("role_id", userRole.ID).
						Where("permission_id", permission.ID).
						Exists()

					if !exists {
						permissionRole := role.RolePermission{
							RoleID:       userRole.ID,
							PermissionID: permission.ID,
						}
						facades.Orm().Query().Create(&permissionRole)
					}
				}
			}
		}
	}


	return nil
}
