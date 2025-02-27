package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type MigrateV20250227095645605CreateTableAdmin struct {
	migrate.MigrateBasic
	upStruct *StructV20250227095645605CreateTableAdminUp
	downStruct *StructV20250227095645605CreateTableAdminDown
}

func NewMigrateV20250227095645605CreateTableAdmin() *MigrateV20250227095645605CreateTableAdmin {
	return &MigrateV20250227095645605CreateTableAdmin{
		upStruct:&StructV20250227095645605CreateTableAdminUp{},
		downStruct:&StructV20250227095645605CreateTableAdminDown{},
	}
}
func (r *MigrateV20250227095645605CreateTableAdmin) Register() {
	r.Tag = "v_2025_02_27_09_56_45_605_create_table_admin"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250227095645605CreateTableAdminUp struct{
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
type StructV20250227095645605CreateTableAdminDown struct{
}
func (*StructV20250227095645605CreateTableAdminUp) TableName() string {
	return "admin"
}
func (*StructV20250227095645605CreateTableAdminDown) TableName() string {
	return "admin"
}
// Up is migration function
func (r *MigrateV20250227095645605CreateTableAdmin) Up(tx *gorm.DB) error{
	
	if err := tx.Migrator().CreateTable(r.upStruct); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250227095645605CreateTableAdmin) Down(tx *gorm.DB) error{
	if err := tx.Migrator().DropTable(r.downStruct); err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250227095645605CreateTableAdmin) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250227095645605CreateTableAdmin) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
