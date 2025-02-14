package tpl

const MigrationCreateTableCode = `package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type Migrate{{TypeTag}} struct {
	migrate.MigrateBasic
	currentTable *Struct{{TypeTag}}
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		currentTable:&Struct{{TypeTag}}{},
	}
}
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
// !!!BEGIN!!!
// Here is the code that you are focusing on

type Struct{{TypeTag}} struct{
	Id        int64 ` + "`" + `gorm:"primaryKey;autoIncrement"` + "`" + `
	CreatedAt int64 ` + "`" + `gorm:"autoCreateTime"` + "`" + `
	UpdatedAt int64 ` + "`" + `gorm:"autoUpdateTime"` + "`" + `
	DeletedAt int64 ` + "`" + `gorm:"index"` + "`" + `
}

func (*Struct{{TypeTag}}) TableName() string {
	return "{{TableName}}"
}
// Up is migration function
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
	
	err := tx.Migrator().CreateTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropTable(r.currentTable)
	if err != nil {
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
	currentTable *Struct{{TypeTag}}
}

func NewMigrate{{TypeTag}}() *Migrate{{TypeTag}} {
	return &Migrate{{TypeTag}}{
		currentTable:&Struct{{TypeTag}}{},
	}
}

// Up is migration function
func (r *Migrate{{TypeTag}}) Register() {
	r.Tag = "{{Tag}}"

}
// !!!BEGIN!!!
// ↓↓↓↓↓↓ Here is the code that you are focusing on

type Struct{{TypeTag}} struct{
	UserName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
	NickName        string ` + "`" + `gorm:"type:varchar(100);"` + "`" + `
}
func (*Struct{{TypeTag}}) TableName() string {
	return "{{TableName}}"
}
func (r *Migrate{{TypeTag}}) Up(tx *gorm.DB) error{
	err := tx.Migrator().AddColumn(r.currentTable,"UserName")
	if err != nil {
			return err
	}
	err = tx.Migrator().AddColumn(r.currentTable,"NickName")
	if err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *Migrate{{TypeTag}}) Down(tx *gorm.DB) error{
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
func (r *Migrate{{TypeTag}}) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *Migrate{{TypeTag}}) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!
`
