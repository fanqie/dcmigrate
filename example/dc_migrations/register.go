package dc_migrations

import (
	"github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
    migrate.RegisterMigration("V20250304153724631CreateTableItems", NewMigrateV20250304153724631CreateTableItems())
	migrate.RegisterMigration("V20250304153732874CreateTableUser", NewMigrateV20250304153732874CreateTableUser())
	migrate.RegisterMigration("V20250304153744225AlterTableItems", NewMigrateV20250304153744225AlterTableItems())
}
