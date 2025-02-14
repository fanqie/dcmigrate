package core

import (
	"database/sql"
	"github.com/fanqie/dcmigrate/pkg/utility"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func MigrateHandle(step int, manage MigratesManage, migrations map[string]DcMigrateImpl) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		utility.InfoPrint("migration start")
		err := manage.RefreshMigrationsData(tx)
		if err != nil {
			return err
		}
		var count int
		for _, m := range manage.PendingList {
			tag := m.GetTypeTag()

			migration := migrations[tag]
			if migration != nil && (step == 0 || count < step) {
				utility.InfoPrintf("migration:%s\n", tag)
				err := migration.Up(tx)
				if err != nil {
					return err
				}
				migration.UpAfter()
				m.SetAlreadyMigrated(true)
				m.SetExecutedAt(sql.NullTime{Time: time.Now(), Valid: true})
				err = tx.Save(&m).Error
				if err != nil {
					return err
				}
				utility.SuccessPrintf("migration count: %s version: %s ok!\n", strconv.Itoa(step), tag)
				count++
			}
		}
		utility.InfoPrintf("migration done, handle count: %s", strconv.Itoa(count))
		return nil

	})

}
