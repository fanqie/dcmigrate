package core

import (
	"github.com/fanqie/dcmigrate/pkg/utility"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DbTool struct {
	Db *gorm.DB
}
type GromParams struct {
	Dialector gorm.Dialector
	Opts      *gorm.Config
}

func Migrator() gorm.Migrator {
	return Db.Migrator()
}
func NewDbTool() *DbTool {
	return &DbTool{}
}
func (r *DbTool) Open(params GromParams) error {
	db, err := gorm.Open(params.Dialector, params.Opts)
	if err != nil {
		utility.ErrPrintf("the database connect error:%s", err.Error())
		return err
	}
	//db = db.Debug()
	Db = db
	r.Db = db
	return err
}
