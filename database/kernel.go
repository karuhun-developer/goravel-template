package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"karuhundeveloper.com/gostarterkit/database/migrations"
	"karuhundeveloper.com/gostarterkit/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateUsersTable{},
		&migrations.M20210101000002CreateJobsTable{},
		&migrations.M20251030060952CreateRolesTable{},
		&migrations.M20251030061125CreatePermissionsTable{},
		&migrations.M20251030061205CreateRolePermissionsTable{},
		&migrations.M20251030064446CreateUserRolesTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.RolePermissionSeeder{},
		&seeders.UserSeeder{},
	}
}
