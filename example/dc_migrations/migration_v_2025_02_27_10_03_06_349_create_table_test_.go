package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type MigrateV20250227100306349CreateTableTest struct {
	migrate.MigrateBasic
	upStruct *StructV20250227100306349CreateTableTestUp
	downStruct *StructV20250227100306349CreateTableTestDown
}

func NewMigrateV20250227100306349CreateTableTest() *MigrateV20250227100306349CreateTableTest {
	return &MigrateV20250227100306349CreateTableTest{
		upStruct:&StructV20250227100306349CreateTableTestUp{},
		downStruct:&StructV20250227100306349CreateTableTestDown{},
	}
}
func (r *MigrateV20250227100306349CreateTableTest) Register() {
	r.Tag = "v_2025_02_27_10_03_06_349_create_table_test"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250227100306349CreateTableTestUp struct{
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
type StructV20250227100306349CreateTableTestDown struct{
}
func (*StructV20250227100306349CreateTableTestUp) TableName() string {
	return "test"
}
func (*StructV20250227100306349CreateTableTestDown) TableName() string {
	return "test"
}
// Up is migration function
func (r *MigrateV20250227100306349CreateTableTest) Up(tx *gorm.DB) error{
	
	if err := tx.Migrator().CreateTable(r.upStruct); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250227100306349CreateTableTest) Down(tx *gorm.DB) error{
	if err := tx.Migrator().DropTable(r.downStruct); err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250227100306349CreateTableTest) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250227100306349CreateTableTest) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
