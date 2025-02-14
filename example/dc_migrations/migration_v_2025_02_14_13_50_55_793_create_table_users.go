package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250214135055793CreateTableUsers struct {
	migrate.MigrateBasic
	currentTable *StructV20250214135055793CreateTableUsers
}

func NewMigrateV20250214135055793CreateTableUsers() *MigrateV20250214135055793CreateTableUsers {
	return &MigrateV20250214135055793CreateTableUsers{
		currentTable:&StructV20250214135055793CreateTableUsers{},
	}
}
func (r *MigrateV20250214135055793CreateTableUsers) Register() {
	r.Tag = "v_2025_02_14_13_50_55_793_create_table_users"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250214135055793CreateTableUsers struct{
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt int64 `gorm:"index"`
}

func (*StructV20250214135055793CreateTableUsers) TableName() string {
	return "users"
}
// Up is migration function
func (r *MigrateV20250214135055793CreateTableUsers) Up(tx *gorm.DB) error{
	
	err := tx.Migrator().CreateTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250214135055793CreateTableUsers) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250214135055793CreateTableUsers) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250214135055793CreateTableUsers) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
