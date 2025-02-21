package core

import (
	"fmt"
	"github.com/fanqie/dcmigrate/pkg/utility"
	"gorm.io/gorm"
)

type MigratesManage struct {
	MigrateList []MigrateBasic
	AlreadyList []MigrateBasic
	PendingList []MigrateBasic
	TableName   string
}

func NewMigratesManage() *MigratesManage {
	return &MigratesManage{}
}
func (r *MigratesManage) RefreshMigrationsData(tx *gorm.DB) error {
	var migrateList []*MigrateBasic
	result := tx.Order("created_at asc").Find(&migrateList)
	if result.Error != nil {
		utility.ErrPrintf("the database connect error:%s", result.Error.Error())
		return result.Error
	}
	r.MigrateList = make([]MigrateBasic, 0)
	r.AlreadyList = make([]MigrateBasic, 0)
	r.PendingList = make([]MigrateBasic, 0)
	for _, s := range migrateList {
		r.MigrateList = append(r.MigrateList, *s)
		if s.AlreadyMigrated {
			r.AlreadyList = append(r.AlreadyList, *s)
		} else {
			r.PendingList = append(r.PendingList, *s)
		}
	}
	for i := len(r.AlreadyList)/2 - 1; i >= 0; i-- {
		opp := len(r.AlreadyList) - 1 - i
		r.AlreadyList[i], r.AlreadyList[opp] = r.AlreadyList[opp], r.AlreadyList[i]
	}
	return nil
}

func (r *MigratesManage) CheckTable(migrations map[string]DcMigrateImpl) {
	utility.InfoPrint("check dc_migrations table")
	if !Db.Migrator().HasTable(&MigrateBasic{}) {
		err := Db.AutoMigrate(&MigrateBasic{})
		if err != nil {
			return
		}
		utility.SuccessPrint("create dc_migrations")
		utility.SuccessPrint("ok!")
	} else {
		utility.InfoPrint("dc_migrations is ok")
	}
	r.checkRegistry(migrations)

}
func (r *MigratesManage) checkRegistry(migrations map[string]DcMigrateImpl) {
	err := Db.Transaction(func(tx *gorm.DB) error {
		fmt.Println(migrations)
		for _, migration := range migrations {
			var count int64
			tx.Model(&MigrateBasic{}).Where("tag = ?", migration.GetTag()).Count(&count)
			fmt.Printf("migration:%v,count:%d\n", migration.GetTag(), count)
			//if count == 0 {
			//	tx.Save(&MigrateBasic{
			//		Tag:             migration.GetTag(),
			//		AlreadyMigrated: false,
			//		CreatedAt:       time.Now(),
			//		UpdatedAt:       time.Now(),
			//	})
			//}
		}
		return nil
	})
	if err != nil {
		utility.ErrPrintf("the database connect error:%s", err.Error())
		panic(err)
	}

}
