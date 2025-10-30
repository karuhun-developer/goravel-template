package role

import (
	"github.com/goravel/framework/database/orm"
)

type Role struct {
	orm.Model
	Name 			string
	RolePermissions []*RolePermission
	Permissions		[]*Permission `gorm:"many2many:role_permissions"`
}
