package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type MigrateV20250220145511357AlterTableUsers struct {
	migrate.MigrateBasic
	upStruct   *StructV20250220145511357AlterTableUsersUp
	downStruct *StructV20250220145511357AlterTableUsersDown
}

func NewMigrateV20250220145511357AlterTableUsers() *MigrateV20250220145511357AlterTableUsers {
	return &MigrateV20250220145511357AlterTableUsers{
		upStruct:   &StructV20250220145511357AlterTableUsersUp{},
		downStruct: &StructV20250220145511357AlterTableUsersDown{},
	}
}

// Up is migration function
func (r *MigrateV20250220145511357AlterTableUsers) Register() {
	r.Tag = "v_2025_02_20_14_55_11_357_alter_table_users"
}

// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type StructV20250220145511357AlterTableUsersUp struct {
	UserName string `gorm:"type:varchar(100);"`
	Id       int32  `gorm:"primaryKey;autoIncrement"`
	NickName string `gorm:"type:varchar(100);"`
}
type StructV20250220145511357AlterTableUsersDown struct {
	Id       uint32 `gorm:"primaryKey;autoIncrement"`
	UserName string `gorm:"type:varchar(100);"`
	NickName string `gorm:"type:varchar(100);"`
}

func (*StructV20250220145511357AlterTableUsersUp) TableName() string {
	return "users"
}
func (*StructV20250220145511357AlterTableUsersDown) TableName() string {
	return "users"
}
func (r *MigrateV20250220145511357AlterTableUsers) Up(tx *gorm.DB) error {
	err := tx.Migrator().AlterColumn(r.upStruct, "Id")
	if err != nil {
		return err
	}
	err = tx.Migrator().AddColumn(r.upStruct, "UserName")
	if err != nil {
		return err
	}
	err = tx.Migrator().AddColumn(r.upStruct, "NickName")
	if err != nil {
		return err
	}
	return nil
}

// Down is rollback function
func (r *MigrateV20250220145511357AlterTableUsers) Down(tx *gorm.DB) error {
	err := tx.Migrator().AlterColumn(r.downStruct, "Id")
	if err != nil {
		return err
	}
	err = tx.Migrator().DropColumn(r.downStruct, "UserName")
	if err != nil {
		return err
	}

	err = tx.Migrator().DropColumn(r.downStruct, "NickName")
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250220145511357AlterTableUsers) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250220145511357AlterTableUsers) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on
// !!!END!!!
