package user

import (
	"github.com/goravel/framework/database/orm"
	"karuhundeveloper.com/gostarterkit/app/models/role"
)

type User struct {
	orm.Model
	Name     string
	Email    string
	Password string
	RoleUser *UserRole
	Roles    []*role.Role `gorm:"many2many:user_roles;joinForeignKey:UserID;JoinReferences:RoleID"`
}
