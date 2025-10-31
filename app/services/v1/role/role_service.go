package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"

	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	models "karuhundeveloper.com/gostarterkit/app/models/role"
)

type RoleService struct {
	// Add your service dependencies here
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (u *RoleService) Create(ctx http.Context, createRequest role.RoleCreateRequest) (roleData models.Role, err error) {
	roleData = models.Role{
		Name: createRequest.Name,
	}

	// Create records
	err = facades.Orm().Query().Create(&roleData)

	return
}

func (u *RoleService) Update(ctx http.Context, updateRequest role.RoleUpdateRequest) (roleData models.Role, err error) {
	// Get records and update
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&roleData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("invalid credentials")
		return
	}

	roleData.Name = updateRequest.Name

	err = facades.Orm().Query().Save(&roleData)

	return
}


func (u *RoleService) Delete(ctx http.Context) (err error) {
	var roleData models.Role

	// Get records and delete
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&roleData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("invalid credentials")
		return
	}

	_, err = facades.Orm().Query().Delete(&roleData)

	return
}