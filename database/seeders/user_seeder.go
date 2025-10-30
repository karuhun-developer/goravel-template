package seeders

import (
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/models/role"
	"karuhundeveloper.com/gostarterkit/app/models/user"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	var superadminUser user.User

	// Password hash
	password, _ := facades.Hash().Make("password")

	// Superadmin email
	email := "superadmin@superadmin.com"

	facades.Orm().Query().Where("email", email).FirstOrCreate(&superadminUser, user.User{
		Name:     "Super Admin",
		Email:    email,
		Password: password,
	})

	// Clear the associated roles first
	facades.Orm().Query().Model(&superadminUser).Association("Roles").Clear()

	// Give the user role of superadmin
	var superadminRole role.Role
	facades.Orm().Query().Where("name", "superadmin").FirstOrCreate(&superadminRole, role.Role{
		Name: "superadmin",
	})

	facades.Orm().Query().Model(&superadminUser).Association("Roles").Append(&superadminRole)

	return nil
}
