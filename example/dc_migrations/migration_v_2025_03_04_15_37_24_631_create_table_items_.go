package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type MigrateV20250304153724631CreateTableItems struct {
	migrate.MigrateBasic
	upStruct *StructV20250304153724631CreateTableItemsUp
	downStruct *StructV20250304153724631CreateTableItemsDown
}

func NewMigrateV20250304153724631CreateTableItems() *MigrateV20250304153724631CreateTableItems {
	return &MigrateV20250304153724631CreateTableItems{
		upStruct:&StructV20250304153724631CreateTableItemsUp{},
		downStruct:&StructV20250304153724631CreateTableItemsDown{},
	}
}
func (r *MigrateV20250304153724631CreateTableItems) Register() {
	r.Tag = "v_2025_03_04_15_37_24_631_create_table_items"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250304153724631CreateTableItemsUp struct{
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
type StructV20250304153724631CreateTableItemsDown struct{
}
func (*StructV20250304153724631CreateTableItemsUp) TableName() string {
	return "items"
}
func (*StructV20250304153724631CreateTableItemsDown) TableName() string {
	return "items"
}
// Up is migration function
func (r *MigrateV20250304153724631CreateTableItems) Up(tx *gorm.DB) error{
	
	if err := tx.Migrator().CreateTable(r.upStruct); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250304153724631CreateTableItems) Down(tx *gorm.DB) error{
	if err := tx.Migrator().DropTable(r.downStruct); err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250304153724631CreateTableItems) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250304153724631CreateTableItems) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
