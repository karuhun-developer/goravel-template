package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251030061125CreatePermissionsTable struct{}

// Signature The unique signature for the migration.
func (r *M20251030061125CreatePermissionsTable) Signature() string {
	return "20251030061125_create_permissions_table"
}

// Up Run the migrations.
func (r *M20251030061125CreatePermissionsTable) Up() error {
	if !facades.Schema().HasTable("permissions") {
		return facades.Schema().Create("permissions", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.Unique("name")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251030061125CreatePermissionsTable) Down() error {
 	return facades.Schema().DropIfExists("permissions")
}
