package user

import (
	"github.com/goravel/framework/database/orm"
	"karuhundeveloper.com/gostarterkit/app/models/role"
)

type UserRole struct {
	orm.Model
	UserId 	uint
	RoleId 	uint
	User  	*User
	Role  	*role.Role
}
