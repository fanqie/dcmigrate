package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250227100348491AlterTableTest struct {
	migrate.MigrateBasic
	upStruct *StructV20250227100348491AlterTableTestUp
	downStruct *StructV20250227100348491AlterTableTestDown
}

func NewMigrateV20250227100348491AlterTableTest() *MigrateV20250227100348491AlterTableTest {
	return &MigrateV20250227100348491AlterTableTest{
		upStruct:&StructV20250227100348491AlterTableTestUp{},
		downStruct:&StructV20250227100348491AlterTableTestDown{},
	}
}

// Up is migration function
func (r *MigrateV20250227100348491AlterTableTest) Register() {
	r.Tag = "v_2025_02_27_10_03_48_491_alter_table_test"
}
// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type StructV20250227100348491AlterTableTestUp struct{
	Id       		int32  `gorm:"primaryKey;autoIncrement"`
	UserName        string `gorm:"type:varchar(100);"`
	NickName        string `gorm:"type:varchar(100);"`
}
type StructV20250227100348491AlterTableTestDown struct{
	Id       		uint32  `gorm:"primaryKey;autoIncrement"`
	UserName        string `gorm:"type:varchar(100);"`
	NickName        string `gorm:"type:varchar(100);"`
}
func (*StructV20250227100348491AlterTableTestUp) TableName() string {
	return "test"
}
func (*StructV20250227100348491AlterTableTestDown) TableName() string {
	return "test"
}
func (r *MigrateV20250227100348491AlterTableTest) Up(tx *gorm.DB) error{
	if err := tx.Migrator().AlterColumn(r.upStruct,"Id");err != nil {
			return err
	}
	if err := tx.Migrator().AddColumn(r.upStruct,"UserName"); err != nil {
			return err
	}
	if err := tx.Migrator().AddColumn(r.upStruct,"NickName"); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250227100348491AlterTableTest) Down(tx *gorm.DB) error{
   if err := tx.Migrator().AlterColumn(r.downStruct, "Id"); err != nil {
		return err
	}
	if err := tx.Migrator().DropColumn(r.downStruct, "UserName"); err != nil {
		return err
	}

	if err := tx.Migrator().DropColumn(r.downStruct, "NickName"); err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250227100348491AlterTableTest) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250227100348491AlterTableTest) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
