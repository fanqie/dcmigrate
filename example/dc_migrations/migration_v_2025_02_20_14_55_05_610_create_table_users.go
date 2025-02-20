package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type MigrateV20250220145505610CreateTableUsers struct {
	migrate.MigrateBasic
	upStruct *StructV20250220145505610CreateTableUsersUp
	downStruct *StructV20250220145505610CreateTableUsersDown
}

func NewMigrateV20250220145505610CreateTableUsers() *MigrateV20250220145505610CreateTableUsers {
	return &MigrateV20250220145505610CreateTableUsers{
		upStruct:&StructV20250220145505610CreateTableUsersUp{},
		downStruct:&StructV20250220145505610CreateTableUsersDown{},
	}
}
func (r *MigrateV20250220145505610CreateTableUsers) Register() {
	r.Tag = "v_2025_02_20_14_55_05_610_create_table_users"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250220145505610CreateTableUsersUp struct{
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
type StructV20250220145505610CreateTableUsersDown struct{
}
func (*StructV20250220145505610CreateTableUsersUp) TableName() string {
	return "users"
}
func (*StructV20250220145505610CreateTableUsersDown) TableName() string {
	return "users"
}
// Up is migration function
func (r *MigrateV20250220145505610CreateTableUsers) Up(tx *gorm.DB) error{
	
	err := tx.Migrator().CreateTable(r.upStruct)
	if err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250220145505610CreateTableUsers) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropTable(r.downStruct)
	if err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250220145505610CreateTableUsers) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250220145505610CreateTableUsers) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
