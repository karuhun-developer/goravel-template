package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"

	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/helpers"
	"karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	"karuhundeveloper.com/gostarterkit/app/http/responses"
	models "karuhundeveloper.com/gostarterkit/app/models/role"
)

type PermissionService struct {
	// Add your service dependencies here
}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

func (u *PermissionService) List(ctx http.Context) (permissionData []models.Permission, pagination responses.PaginationResponse, err error) {
	// Set default pagination values
	page := ctx.Request().QueryInt("page", 1)
	paginate := ctx.Request().QueryInt("paginate", 10)

	// Get records with pagination
	var total int64

	query := facades.Orm().Query()

	// Apply filters
	fields := []string{
		"name",
	}
	// Filter helper
	query = helpers.OrmFilter(ctx, query, fields)

	err = query.Paginate(page, paginate, &permissionData, &total)

	// Return if error occurs
	if err != nil {
		return
	}

	pagination, err = helpers.PaginateHelper(page, paginate, total, &permissionData)

	return
}

func (u *PermissionService) Show(ctx http.Context) (permissionData models.Permission, err error) {
	// Get records
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&permissionData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("Permission not found")
		return
	}

	return
}

func (u *PermissionService) Create(ctx http.Context, createRequest role.PermissionCreateRequest) (permissionData models.Permission, err error) {
	permissionData = models.Permission{
		Name: createRequest.Name,
	}

	// Create records
	err = facades.Orm().Query().Create(&permissionData)

	return
}

func (u *PermissionService) Update(ctx http.Context, updateRequest role.PermissionUpdateRequest) (permissionData models.Permission, err error) {
	// Get records and update
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&permissionData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("Permission not found")
		return
	}

	permissionData.Name = updateRequest.Name

	err = facades.Orm().Query().Save(&permissionData)

	return
}


func (u *PermissionService) Delete(ctx http.Context) (err error) {
	var permissionData models.Permission

	// Get records and delete
	err = facades.Orm().Query().
		Where("id", ctx.Request().Route("id")).
		FirstOrFail(&permissionData)

	// Return if record not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("Permission not found")
		return
	}

	_, err = facades.Orm().Query().Delete(&permissionData)

	return
}