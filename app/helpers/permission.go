package helpers

import (
	"time"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/str"
	"karuhundeveloper.com/gostarterkit/app/models/role"
	"karuhundeveloper.com/gostarterkit/app/models/user"
)

func getCachePermissionPrefix() string {
	appName := facades.Config().Env("APP_NAME", "Goravel").(string)
	return str.Of(appName).Kebab().String()
}

func GetUserRole(authUser user.User) (userRole user.UserRole, err error) {
	// Query the user role
	err = facades.Orm().Query().Model(&user.UserRole{}).
		Join("join users on users.id = user_roles.user_id").
		Where("users.id", authUser.ID).
		Select("user_roles.*").
		First(&userRole)

	return
}

func HasRole(authUser user.User, roleName string) bool {
	cacheKey := getCachePermissionPrefix() + "role:" + roleName + ":" + string(rune(authUser.ID))

	// Check cache first
	if facades.Cache().Has(cacheKey) {
		return facades.Cache().Get(cacheKey).(bool)
	}

	// Query the user role
	exists, err := facades.Orm().Query().Model(&user.UserRole{}).
		Join("join users on users.id = user_roles.user_id").
		Join("join roles on roles.id = user_roles.role_id").
		Where("users.id", authUser.ID).
		Where("roles.name", roleName).
		Exists()

	// Return true if the role exists for the user and cache the result
	if err == nil && exists {
		facades.Cache().Add(cacheKey, true, 1*time.Hour)
		return true
	}

	// Cache negative result to avoid repeated DB hits (optional: short TTL)
	facades.Cache().Add(cacheKey, false, 1*time.Hour)

	return false
}

func HasPermission(authUser user.User, permissionName string) bool {
	hasCachePermission := facades.Cache().Has(getCachePermissionPrefix() + permissionName + ":" + string(rune(authUser.ID)))

	// Check cache first
	if hasCachePermission {
		return facades.Cache().Get(getCachePermissionPrefix() + permissionName + ":" + string(rune(authUser.ID))).(bool)
	}

	var RolePermission role.RolePermission

	// Query the user role
	err := facades.Orm().Query().Model(&role.RolePermission{}).
		Join("join roles on roles.id = role_permissions.role_id").
		Join("join user_roles on user_roles.role_id = roles.id").
		Join("join users on users.id = user_roles.user_id").
		Where("users.id", authUser.ID).
		Where("role_permissions.name", permissionName).
		Select("role_permissions.*").
		FirstOrFail(&RolePermission)

	// Return true if the permission exists for the user
	if err == nil {
		// Save to cache for 1 hour
		success := facades.Cache().Add(getCachePermissionPrefix()+permissionName+":"+string(rune(RolePermission.RoleID)), true, 1*time.Hour)
		return success
	}

	// Cache negative result to avoid repeated DB hits (optional: short TTL)
	facades.Cache().Add(getCachePermissionPrefix()+permissionName+":"+string(rune(authUser.ID)), false, 1*time.Hour)

	return false
}