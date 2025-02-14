package pkg

import (
	"fmt"
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
	r.MigrationsManage = core.NewMigratesManage()
	r.MigrationsManage.CheckTable()
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
