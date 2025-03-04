package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250304153744225AlterTableItems struct {
	migrate.MigrateBasic
	upStruct *StructV20250304153744225AlterTableItemsUp
	downStruct *StructV20250304153744225AlterTableItemsDown
}

func NewMigrateV20250304153744225AlterTableItems() *MigrateV20250304153744225AlterTableItems {
	return &MigrateV20250304153744225AlterTableItems{
		upStruct:&StructV20250304153744225AlterTableItemsUp{},
		downStruct:&StructV20250304153744225AlterTableItemsDown{},
	}
}

// Up is migration function
func (r *MigrateV20250304153744225AlterTableItems) Register() {
	r.Tag = "v_2025_03_04_15_37_44_225_alter_table_items"
}
// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type StructV20250304153744225AlterTableItemsUp struct{
	Id       		int32  `gorm:"primaryKey;autoIncrement"`
	UserName        string `gorm:"type:varchar(100);"`
	NickName        string `gorm:"type:varchar(100);"`
}
type StructV20250304153744225AlterTableItemsDown struct{
	Id       		uint32  `gorm:"primaryKey;autoIncrement"`
	UserName        string `gorm:"type:varchar(100);"`
	NickName        string `gorm:"type:varchar(100);"`
}
func (*StructV20250304153744225AlterTableItemsUp) TableName() string {
	return "items"
}
func (*StructV20250304153744225AlterTableItemsDown) TableName() string {
	return "items"
}
func (r *MigrateV20250304153744225AlterTableItems) Up(tx *gorm.DB) error{
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
func (r *MigrateV20250304153744225AlterTableItems) Down(tx *gorm.DB) error{
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
func (r *MigrateV20250304153744225AlterTableItems) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250304153744225AlterTableItems) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
