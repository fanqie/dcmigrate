## Edit the migration file
open "./dc_migrations/migration_${Tag}.go" file and edit it
```golang
package dc_migrations

import (
	migrate "github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/gorm"
)

type Migrate_Tag_CreateTableUsers struct {
	migrate.MigrateBasic
	currentTable *Struct_Tag_CreateTableUsers
}

func NewMigrate_Tag_CreateTableUsers() *Migrate_Tag_CreateTableUsers {
	return &Migrate_Tag_CreateTableUsers{
		currentTable:&Struct_Tag_CreateTableUsers{},
	}
}
func (r *Migrate_Tag_CreateTableUsers) Register() {
	r.Tag = "v_2025_02_14_09_48_00_702_create_table_users"
}
// !!!BEGIN!!!
// Here is the code that you are focusing on
// You can add your code here

// this is the struct that you want to create
type Struct_Tag_CreateTableUsers struct{
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt int64 `gorm:"index"`
}

func (*Struct_Tag_CreateTableUsers) TableName() string {
	return "users" //Can be modified to your target table name
}
// Up is migration function
func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{

	err := tx.Migrator().CreateTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
// Down is rollback function
func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
	err := tx.Migrator().DropTable(r.currentTable)
	if err != nil {
		return err
	}
	return nil
}
func (r *Migrate_Tag_CreateTableUsers) AfterUp(tx *gorm.DB) {
	//run in after "Up function"
}
func (r *Migrate_Tag_CreateTableUsers) AfterDown(tx *gorm.DB) {
	//run in after "Down function"
}

// ↑↑↑↑↑↑ Here is the code that you are focusing on 
// !!!END!!!

```
###  Write migration func
####  Create table
1. generate migration file
    ```shell
    $ go run dmc.go gen --create users
    ```
    ```shell
    [Info]check dc_migrations table
    [Success]create dc_migrations
    [Success]ok!
    [Info]create migration start
    [Success]ok! file name :[./dc_migrations/migration_v_2025_02_14_09_48_00_702_create_table_users.go]
    ```
1. define the struct
    ```go
    type Struct_Tag_CreateTableUsers struct{
        Id        int64 `gorm:"primaryKey;autoIncrement"`
        CreatedAt int64 `gorm:"autoCreateTime"`
        UpdatedAt int64 `gorm:"autoUpdateTime"`
        DeletedAt int64 `gorm:"index"`
    }
    ```
1. write your migration code
    ```go
    // Up is migration function
    func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().CreateTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
1. Write your rollback func
    ```go
    // Down is rollback function
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().DropTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
####  Alter column
1. generate migration file
    ```shell
    $ go run dmc.go gen --alter users
    ```
    ```shell
    # output
    [Info]check dc_migrations table
    [Info]dc_migrations is ok
    [Info]create migration start
    [Success]ok! file name :[./dc_migrations/migration_v_2025_02_14_09_55_03_505_alter_table_users.go]
    ```
1. define the struct
    ```go
    type Struct_Tag_CreateTableUsers struct{
        Id        int64 `gorm:"primaryKey;autoIncrement"`
        CreatedAt int64 `gorm:"autoCreateTime"`
        UpdatedAt int64 `gorm:"autoUpdateTime"`
        DeletedAt int64 `gorm:"index"`
    }
    ```
1. write your migration code
    ```go
    // Up is migration function
    func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().CreateTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
1. Write your rollback func
    ```go
    // Down is rollback function
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().DropTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
###  Alter table name
Just modify the return value
```go
func (*Struct_Tag_CreateTableUsers) TableName() string {
	return "users" //Can be modified to your target table name
}
```

###  [Gorm Fields Tags](https://gorm.io/docs/models.html#Fields-Tags)
