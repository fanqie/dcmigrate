package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250214140556456AlterTableUsers struct {
	migrate.MigrateBasic
	currentTable *StructV20250214140556456AlterTableUsers
}

func NewMigrateV20250214140556456AlterTableUsers() *MigrateV20250214140556456AlterTableUsers {
	return &MigrateV20250214140556456AlterTableUsers{
		currentTable: &StructV20250214140556456AlterTableUsers{},
	}
}

// Up is migration function
func (r *MigrateV20250214140556456AlterTableUsers) Register() {
	r.Tag = "v_2025_02_14_14_05_56_456_alter_table_users"

}

// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type StructV20250214140556456AlterTableUsers struct {
	UserName string `gorm:"type:varchar(100);"`
	NickName string `gorm:"type:varchar(100);"`
}

func (*StructV20250214140556456AlterTableUsers) TableName() string {
	return "users"
}
func (r *MigrateV20250214140556456AlterTableUsers) Up(tx *gorm.DB) error {
	err := tx.Migrator().AddColumn(r.currentTable, "UserName")
	if err != nil {
		return err
	}
	err = tx.Migrator().AddColumn(r.currentTable, "NickName")
	if err != nil {
		return err
	}
	return nil
}

// Down is rollback function
func (r *MigrateV20250214140556456AlterTableUsers) Down(tx *gorm.DB) error {
	err := tx.Migrator().DropColumn(r.currentTable, "UserName")
	if err != nil {
		return err
	}

	err = tx.Migrator().DropColumn(r.currentTable, "NickName")
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250214140556456AlterTableUsers) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250214140556456AlterTableUsers) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on
// !!!END!!!
