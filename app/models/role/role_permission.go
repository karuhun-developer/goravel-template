package role

import (
	"github.com/goravel/framework/database/orm"
)

type RolePermission struct {
	orm.Model
	RoleID       	uint
	PermissionID 	uint
	Role 			*Role
	Permission		*Permission
}