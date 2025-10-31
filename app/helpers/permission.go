package helpers

import (
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/str"
	"karuhundeveloper.com/gostarterkit/app/models/role"
	"karuhundeveloper.com/gostarterkit/app/models/user"
)

func getCachePermissionPrefix() string {
	appName := facades.Config().Env("APP_NAME", "Goravel").(string) + ":"
	return str.Of(appName).Kebab().String()
}

/**
 GetUserRole retrieves the role of the currently authenticated user. It first checks the cache for the user's role.
 If not found in the cache, it queries the database and stores the result in the cache for future requests.

 Parameters:
 - ctx: The HTTP context containing request information.

 Returns:
 - userRole: The UserRole model of the authenticated user.
 - err: An error object if any issues occur during retrieval.
*/
func GetUserRole(ctx http.Context) (userRole user.UserRole, err error) {
	var authUser user.User

	err = facades.Auth(ctx).User(&authUser)

	if err != nil {
		return
	}

	// Generate cache key
	cacheKey := getCachePermissionPrefix() + "user:role:" + strconv.Itoa(int(authUser.ID))

	// Check cache first
	if facades.Cache().Has(cacheKey) {
		return facades.Cache().Get(cacheKey).(user.UserRole), nil
	}

	// Query the user role
	err = facades.Orm().Query().Model(&user.UserRole{}).
		Join("join users on users.id = user_roles.user_id").
		Where("users.id", authUser.ID).
		Select("user_roles.*").
		First(&userRole)

	if err == nil {
		return
	}

	// Save to cache for 1 hour
	facades.Cache().Add(cacheKey, userRole, 1*time.Hour)

	return
}

/**
 HasRole checks if the currently authenticated user has a specific role. The result is cached to optimize performance for subsequent requests.

 Parameters:
 - ctx: The HTTP context containing request information.
 - roleName: The name of the role to check.

 Returns:
 - bool: True if the user has the specified role, false otherwise.
*/
func HasRole(ctx http.Context, roleName string) bool {
	userRole, err := GetUserRole(ctx)

	// Return false if unable to get user role
	if err != nil {
		return false
	}

	// Generate cache key
	cacheKey := getCachePermissionPrefix() + "role:" + roleName + ":" + strconv.Itoa(int(userRole.UserId))

	// Check cache first
	if facades.Cache().Has(cacheKey) {
		return StringToBool(facades.Cache().Get(cacheKey).(string))
	}

	var Role role.Role

	// Query the user role
	err = facades.Orm().Query().Model(&role.Role{}).
		Join("join user_roles on user_roles.role_id = roles.id").
		Join("join users on users.id = user_roles.user_id").
		Where("users.id", userRole.UserId).
		Where("roles.name", roleName).
		Select("roles.*").
		FirstOrFail(&Role)

	// If role not found, cache negative result
	if errors.Is(err, errors.OrmRecordNotFound) {
		// Cache negative result to avoid repeated DB hits (optional: short TTL)
		facades.Cache().Add(cacheKey, false, 5*time.Minute)

		return false
	}

	// Save to cache for 1 hour
	success := facades.Cache().Add(cacheKey, true, 1*time.Hour)

	return success
}

/**
 HasPermission checks if the currently authenticated user has a specific permission. It first retrieves the user's role and then checks if the permission exists for that role.
 The result is cached to optimize performance for subsequent requests.

 Parameters:
 - ctx: The HTTP context containing request information.
 - permissionName: The name of the permission to check.

 Returns:
 - bool: True if the user has the specified permission, false otherwise.
*/
func HasPermission(ctx http.Context, permissionName string) bool {
	userRole, err := GetUserRole(ctx)

	// Return false if unable to get user role
	if err != nil {
		return false
	}

	// Generate cache key
	cacheKey := getCachePermissionPrefix() + "permission:" + permissionName + ":" + strconv.Itoa(int(userRole.RoleId))

	facades.Cache().Flush()
	// Check cache first
	if facades.Cache().Has(cacheKey) {
		return StringToBool(facades.Cache().Get(cacheKey).(string))
	}

	var RolePermission role.RolePermission

	// Query the user role
	err = facades.Orm().Query().Model(&role.RolePermission{}).
		Join("join roles on roles.id = role_permissions.role_id").
		Join("join permissions on permissions.id = role_permissions.permission_id").
		Where("roles.id", userRole.RoleId).
		Where("permissions.name", permissionName).
		Select("role_permissions.*").
		FirstOrFail(&RolePermission)

	// If permission not found, cache negative result
	if errors.Is(err, errors.OrmRecordNotFound) {
		println(false)
		// Cache negative result to avoid repeated DB hits (optional: short TTL)
		facades.Cache().Add(cacheKey, false, 5*time.Minute)

		return false
	}

	// Save to cache for 1 hour
	success := facades.Cache().Add(cacheKey, true, 1*time.Hour)

	return success
}

/**
 NoPermissionError creates a standardized error indicating that the user does not have the required permission to perform an action.

 Returns:
 - error: An error object with a predefined message.
*/
func NoPermissionError() error {
	return errors.New("you do not have permission to perform this action")
}