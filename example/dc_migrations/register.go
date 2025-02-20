package dc_migrations

import (
	"github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
    migrate.RegisterMigration("V20250220145505610CreateTableUsers", NewMigrateV20250220145505610CreateTableUsers())
	migrate.RegisterMigration("V20250220145511357AlterTableUsers", NewMigrateV20250220145511357AlterTableUsers())
}
