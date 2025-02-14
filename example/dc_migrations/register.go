package dc_migrations

import (
	"github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
    migrate.RegisterMigration("V20250214135055793CreateTableUsers", NewMigrateV20250214135055793CreateTableUsers())
	migrate.RegisterMigration("V20250214140556456AlterTableUsers", NewMigrateV20250214140556456AlterTableUsers())
}
