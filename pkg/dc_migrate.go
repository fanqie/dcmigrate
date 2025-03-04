package pkg

import (
	"fmt"
	"os"

	"github.com/fanqie/dcmigrate/pkg/core"
)

type DcMigrate struct {
	migrations       map[string]core.DcMigrateImpl
	DbTool           *core.DbTool
	MigrationsManage *core.MigratesManage
	isDebug          bool
}

func (r *DcMigrate) RegisterMigration(name string, migrationFunc core.DcMigrateImpl) {
	r.migrations[name] = migrationFunc
}
func NewDcMigrate(isDebug bool) *DcMigrate {
	boot := &DcMigrate{
		migrations: make(map[string]core.DcMigrateImpl),
		DbTool:     core.NewDbTool(),
		isDebug:    isDebug,
	}

	return boot
}

func (r *DcMigrate) Setup(db core.GromParams, afterHandle func()) {
	r.databaseInit(db)
	if r.isDebug {
		r.DbTool.Db.Debug()
	}
	for _, migration := range r.migrations {
		migration.Register()
	}
	r.MigrationsManage = core.NewMigratesManage()
	r.MigrationsManage.CheckTable(r.migrations)
	os.Args = os.Args[1:]
	core.DefinedCommand(r.MigrationsManage, r.migrations)
	afterHandle()

}
func (r *DcMigrate) databaseInit(db core.GromParams) {
	err := r.DbTool.Open(db)
	if err != nil {
		fmt.Println(err)
		panic("the database connect error")
	}

}
