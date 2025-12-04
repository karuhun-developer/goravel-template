package role

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"

	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	models "karuhundeveloper.com/gostarterkit/app/models/role"
)

type RolePermissionService struct {
	// Add your service dependencies here
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{}
}

func (u *RolePermissionService) Show(ctx http.Context) (roleData models.Role, err error) {
	// Get records
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		With("RolePermissions.Permission").
		FirstOrFail(&roleData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("Role not found")
		return
	}

	return
}

func (u *RolePermissionService) Sync(ctx http.Context, syncRequest role.RolePermssionSyncRequest) (role models.Role, err error) {
	// Get records
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&role)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("Role not found")
		return
	}

	// Sync permissions
	tx, _ := facades.DB().BeginTransaction()
	facades.Orm().Query().Model(&models.RolePermission{}).Where("role_id", role.ID).Delete(&models.RolePermission{})

	for _, permID := range ctx.Request().All()["permissions"].([]any) {
		// Check if permission exists
		var permission models.Permission
		err = facades.Orm().Query().
			Where("id", permID).
			FirstOrFail(&permission)

		if (errors.Is(err, errors.OrmRecordNotFound)) {
			err = errors.New(fmt.Sprintf("Permission with ID %v not found", permID))

			tx.Rollback()

			return
		}

		// Check if association already exists
		exists, _ := facades.Orm().Query().Model(&models.RolePermission{}).
			Where("role_id", role.ID).
			Where("permission_id", permission.ID).
			Exists()

		if !exists {
			permissionRole := models.RolePermission{
				RoleID:       role.ID,
				PermissionID: permission.ID,
			}
			facades.Orm().Query().Create(&permissionRole)
		}
	}

	tx.Commit()

	return
}