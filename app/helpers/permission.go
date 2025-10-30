package helpers

import (
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/models/role"
	"karuhundeveloper.com/gostarterkit/app/models/user"
)

const cachePermissionPrefix = "gostarterkit:permission:"

func GetUserRole(authUser user.User) (userRole user.UserRole, err error) {
	// Query the user role
	err = facades.Orm().Query().Model(&user.UserRole{}).Join("join users on users.id = user_roles.user_id").Where("users.id", authUser.ID).Select("user_roles.*").First(&userRole)

	return
}

func HasRole(authUser user.User, roleName string) bool {
	// Query the user role
	exists, err := facades.Orm().Query().Model(&user.UserRole{}).Join("join users on users.id = user_roles.user_id").Join("join roles on roles.id = user_roles.role_id").Where("users.id", authUser.ID).Where("roles.name", roleName).Exists()

	// Return true if the role exists for the user
	if err == nil && exists {
		return true
	}

	return false
}

func HasPermission(authUser user.User, permissionName string) bool {
	var RolePermission role.RolePermission

	// Query the user role
	err := facades.Orm().Query().Model(&role.RolePermission{}).Join("join roles on roles.id = role_permissions.role_id").Join("join user_roles on user_roles.role_id = roles.id").Join("join users on users.id = user_roles.user_id").Where("users.id", authUser.ID).Where("role_permissions.name", permissionName).Select("role_permissions.*").FirstOrFail(&RolePermission)

	// Save to cache
	success := facades.Cache().Add(cachePermissionPrefix+permissionName+":"+string(rune(RolePermission.RoleID)), true, 0)

	// Return true if the permission exists for the user
	if err == nil && success {
		return true
	}

	return false
}