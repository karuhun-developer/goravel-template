package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251030064446CreateUserRolesTable struct{}

// Signature The unique signature for the migration.
func (r *M20251030064446CreateUserRolesTable) Signature() string {
	return "20251030064446_create_user_roles_table"
}

// Up Run the migrations.
func (r *M20251030064446CreateUserRolesTable) Up() error {
	if !facades.Schema().HasTable("user_roles") {
		return facades.Schema().Create("user_roles", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("user_id")
			table.UnsignedBigInteger("role_id")
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.Foreign("role_id").References("id").On("roles").CascadeOnDelete()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251030064446CreateUserRolesTable) Down() error {
 	return facades.Schema().DropIfExists("user_roles")
}
