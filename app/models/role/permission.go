package role

import "github.com/goravel/framework/database/orm"

type Permission struct {
	orm.Model
	Name string
	RolePermissions []*RolePermission
	Roles			[]*Role `gorm:"many2many:role_permissions"`
}
