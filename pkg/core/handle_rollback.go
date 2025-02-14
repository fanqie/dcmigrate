package core

import (
	"database/sql"
	"github.com/fanqie/dcmigrate/pkg/utility"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func RollbackHandle(step int, manage MigratesManage, migrations map[string]DcMigrateImpl) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		utility.InfoPrint("rollback start")
		err := manage.RefreshMigrationsData(tx)
		if err != nil {
			return err
		}
		var count int = 0

		for _, m := range manage.AlreadyList {
			tag := m.GetTypeTag()
			migration := migrations[tag]
			if migration != nil && (step == 0 || count < step) {
				utility.InfoPrintf("rollback(%s):%s\n", strconv.Itoa(step), tag)
				err := migration.Down(tx)
				if err != nil {
					return err
				}
				migration.UpAfter()
				m.SetAlreadyMigrated(false)
				m.SetRevertedAt(sql.NullTime{Time: time.Now(), Valid: true})
				err = tx.Save(&m).Error
				if err != nil {
					return err
				}
				utility.SuccessPrintf("rollback count:%s version: %s ok!\n", strconv.Itoa(step), tag)
				count++

			}

		}
		utility.InfoPrintf("rollback done, handle count: %s", strconv.Itoa(count))
		return nil

	})

}
