package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251030061205CreateRolePermissionsTable struct{}

// Signature The unique signature for the migration.
func (r *M20251030061205CreateRolePermissionsTable) Signature() string {
	return "20251030061205_create_role_permissions_table"
}

// Up Run the migrations.
func (r *M20251030061205CreateRolePermissionsTable) Up() error {
	if !facades.Schema().HasTable("role_permissions") {
		return facades.Schema().Create("role_permissions", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("role_id")
			table.UnsignedBigInteger("permission_id")
			table.Foreign("role_id").References("id").On("roles").CascadeOnDelete()
			table.Foreign("permission_id").References("id").On("permissions").CascadeOnDelete()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251030061205CreateRolePermissionsTable) Down() error {
 	return facades.Schema().DropIfExists("role_permissions")
}
