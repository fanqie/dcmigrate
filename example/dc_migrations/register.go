package dc_migrations

import (
	"github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
    migrate.RegisterMigration("V20250220145505610CreateTableUsers", NewMigrateV20250220145505610CreateTableUsers())
	migrate.RegisterMigration("V20250220145511357AlterTableUsers", NewMigrateV20250220145511357AlterTableUsers())
	migrate.RegisterMigration("V20250227095645605CreateTableAdmin", NewMigrateV20250227095645605CreateTableAdmin())
	migrate.RegisterMigration("V20250227100306349CreateTableTest", NewMigrateV20250227100306349CreateTableTest())
	migrate.RegisterMigration("V20250227100348491AlterTableTest", NewMigrateV20250227100348491AlterTableTest())
}
