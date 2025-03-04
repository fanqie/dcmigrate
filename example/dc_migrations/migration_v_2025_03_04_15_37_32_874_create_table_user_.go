package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type MigrateV20250304153732874CreateTableUser struct {
	migrate.MigrateBasic
	upStruct *StructV20250304153732874CreateTableUserUp
	downStruct *StructV20250304153732874CreateTableUserDown
}

func NewMigrateV20250304153732874CreateTableUser() *MigrateV20250304153732874CreateTableUser {
	return &MigrateV20250304153732874CreateTableUser{
		upStruct:&StructV20250304153732874CreateTableUserUp{},
		downStruct:&StructV20250304153732874CreateTableUserDown{},
	}
}
func (r *MigrateV20250304153732874CreateTableUser) Register() {
	r.Tag = "v_2025_03_04_15_37_32_874_create_table_user"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type StructV20250304153732874CreateTableUserUp struct{
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
type StructV20250304153732874CreateTableUserDown struct{
}
func (*StructV20250304153732874CreateTableUserUp) TableName() string {
	return "user"
}
func (*StructV20250304153732874CreateTableUserDown) TableName() string {
	return "user"
}
// Up is migration function
func (r *MigrateV20250304153732874CreateTableUser) Up(tx *gorm.DB) error{
	
	if err := tx.Migrator().CreateTable(r.upStruct); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *MigrateV20250304153732874CreateTableUser) Down(tx *gorm.DB) error{
	if err := tx.Migrator().DropTable(r.downStruct); err != nil {
		return err
	}
	return nil
}
func (r *MigrateV20250304153732874CreateTableUser) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *MigrateV20250304153732874CreateTableUser) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
