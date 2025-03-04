package tpl

const MigrationCreateTableCode = `package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
	"time"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
	upStruct *Struct{{TypeTag}}Up
	downStruct *Struct{{TypeTag}}Down
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		upStruct:&Struct{{TypeTag}}Up{},
		downStruct:&Struct{{TypeTag}}Down{},
	}
}
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type Struct{{TypeTag}}Up struct{
	Id        uint32 ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	CreatedAt time.Time ` + "`" + `gorm:"autoCreateTime"` + "`" + `
	UpdatedAt time.Time ` + "`" + `gorm:"autoUpdateTime"` + "`" + `
	DeletedAt time.Time ` + "`" + `gorm:"index"` + "`" + `
}
type Struct{{TypeTag}}Down struct{
}
func (*Struct{{TypeTag}}Up) TableName() string {
	return "{{TableName}}"
}
func (*Struct{{TypeTag}}Down) TableName() string {
	return "{{TableName}}"
}
// Up is migration function
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
	
	if err := tx.Migrator().CreateTable(r.upStruct); err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
	if err := tx.Migrator().DropTable(r.downStruct); err != nil {
		return err
	}
	return nil
}
func (r *Migrate{{TypeTag}}) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *Migrate{{TypeTag}}) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
`

const MigrationAlterTableCode = `package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
	upStruct *Struct{{TypeTag}}Up
	downStruct *Struct{{TypeTag}}Down
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		upStruct:&Struct{{TypeTag}}Up{},
		downStruct:&Struct{{TypeTag}}Down{},
	}
}

// Up is migration function
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"
}
// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type Struct{{TypeTag}}Up struct{
	Id       		int32  ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	UserName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
	NickName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
}
type Struct{{TypeTag}}Down struct{
	Id       		uint32  ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	UserName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
	NickName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
}
func (*Struct{{TypeTag}}Up) TableName() string {
	return "{{TableName}}"
}
func (*Struct{{TypeTag}}Down) TableName() string {
	return "{{TableName}}"
}
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
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
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
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
func (r *Migrate{{TypeTag}}) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *Migrate{{TypeTag}}) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
`
