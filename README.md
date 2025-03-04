**English | [中文](doc/zh_cn/README.md)**

Dcmigrate is a database migration tool based on Gorm, which complements the Gorm migration mechanism without any functional conflicts, aiming to provide a more user-friendly database migration tool

# Features
- Supports mysql, sqlite, postgres, sqlserver, tidb, clickhouse and other databases
- Command line tool
- Depends on gorm
- Supports executing migration files in steps
- Supports rollback function
- Can view migration list and status
- Explicit "up" and "down" function implementation
- Supports repairing migration records

# Tooltip ⚠

- **The current version is relatively stable and will continue to be updated. It is recommended to keep it until the version is updated**


# Quick start
## Install dependencies
```shell
go get -u github.com/fanqie/dcmigrate
```
## Quickly Initialize Your Project
[Guide Doc](doc/en/Initialize_Guide.md)


## Directory Structure
This is the directory structure you obtained after initializing the project
```shell
example/
├── dc_migrations // This is the migration file directory for dcmigration
  └── register.go // This is the migration file registration file for dcmigration, which is automatically generated and maintained by dcmigration. Please do not manually modify it
  └── 20230301_000000_create_users_table.go // This is the migration file for dcmigration
  └── dmc.go // This is the command-line tool for dcmigration
├── go.mod
├── go.sum
└── ... you project files

```
## Command line usage
**You must complete the initialization operation**
```shell
go run . dmc --help
```
```shell
[Info]check dc_migrations table
[Info]dc_migrations is ok
Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         generate a new core file
        syntax：dmc gen  [--create|--alter]  {table_name}
        usage：`dmc gen --create user` //or `dmc gen --alter user`
  help        Help about any command
  list        show all migrations record
  migrate     all new migration file versions will be migrated or target step size version
  rollback    rollback history migrates
        syntax：dmc rollback  [--step=1] [--all]  {table_name}
        usage：`dmc rollback --step=2` //or `dmc rollback --all`

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.

```

## Connect to Database
Open the "dc_migrations/dmc.go" file, modify the database connection information, and then run the dmc. go file. You can configure the database connection according to the reference code and the official Gorm documentation

**[Gorm Connecting Database Doc Guide](https://gorm.io/docs/connecting_to_the_database.html)**
```shell
package dc_migrations

import (
	"fmt"
	"os"

	"github.com/fanqie/dcmigrate/pkg"
	"github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TryStartUpDcMigrate() bool {
	if len(os.Args) <= 1 || os.Args[1] != "dmc" {
		return false
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("dcmigrate error, %v", err)
		}
	}()
	dcMigrate := pkg.NewDcMigrate(true)
	Register(dcMigrate)
	dcMigrate.Setup(core.GromParams{
		Dialector: mysqlDialector(),
		// or ↓↓↓↓↓↓↓↓↓↓
		//Dialector: sqliteDialector(),
		//Dialector: pgsqlDialector(),
		//Dialector: sqlserverDialector(),
		//Dialector: tiDBDialector(),
		//Dialector: clickhouseDialector(),
		Opts: &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	}, func() {

	})
	return true
}

// connecting_to_the_database more doc:  https://gorm.io/docs/connecting_to_the_database.html

func mysqlDialector() gorm.Dialector {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	return mysql.Open(dsn)
}

// go get -u gorm.io/driver/sqlite

//func sqliteDialector() gorm.Dialector {
//	dsn := "test.db"
//	return sqlite.Open(dsn)
//}

// go get -u gorm.io/driver/postgres

//func pgsqlDialector() gorm.Dialector {
//	dsn := "user=root password=root dbname=test port=6943 sslmode=disable TimeZone=Local"
//	return postgres.Open(dsn)
//}

// sqlserver
// go get -u gorm.io/driver/sqlserver

//func sqlserverDialector() gorm.Dialector {
//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
//	return sqlserver.Open(dsn)
//}

// TiDB
// go get -u gorm.io/driver/mysql

//func tiDBDialector() gorm.Dialector {
//	dsn := "root:@tcp(127.0.0.1:4000)/test"
//	return mysql.Open(dsn)
//}

// Clickhouse
// go get -u gorm.io/driver/clickhouse

//func clickhouseDialector() gorm.Dialector {
//	dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
//	return clickhouse.Open(dsn)
//}

```
## Start Up DcMigrate In Your Project
```shell
package main
import "dc_migrations"
func main() {
	if dc_migrations.TryStartUpDcMigrate() {
		return
	}
	//todo: Your code here...
	// gin.Run()
}
```
## Generate
### Generate a createTable migration file
```shell
go run . dmc gen --create users
```
```shell
[Info]check dc_migrations table
[Success]create dc_migrations
[Success]ok!
[Info]create migration start
[Success]ok! file name :[./dc_migrations/migration_v_2025_02_14_09_48_00_702_create_table_users.go]
```
### Edit the migration file
[Guide Doc](doc/en/Edit_Migration.md)
### Generate a alterTable migration file
```shell
go run . dmc gen --alter users
```
```shell
# output
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]create migration start
[Success]ok! file name :[./dc_migrations/migration_v_2025_02_14_09_55_03_505_alter_table_users.go]
```
## Show Migrations List
```shell
go run . dmc list
```
```shell
[Info]check dc_migrations table
[Info]dc_migrations is ok
┌────┬──────────────────────────────────────────────┬─────────────────────┬─────────────────┐
│ ID │ TAG                                          │ CREATEDAT           │ ALREADYMIGRATED │
├────┼──────────────────────────────────────────────┼─────────────────────┼─────────────────┤
│  3 │ v_2025_02_14_13_50_55_793_create_table_users │ 2025-02-14 13:50:55 │ ☑ Yes!          │
│  5 │ v_2025_02_14_14_05_56_456_alter_table_users  │ 2025-02-14 14:05:56 │ Pending         │
└────┴──────────────────────────────────────────────┴─────────────────────┴─────────────────┘
```

## Run Migration
### Migrate All
```shell
go run . dmc migrate       
```
```shell
# output
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]migration start
[Info]migration:V20250214094800702CreateTableUsers
[Success]migration count: 0 version: V20250214094800702CreateTableUsers ok!
[Info]migration:V20250214095503505AlterTableUsers
[Success]migration count: 0 version: V20250214095503505AlterTableUsers ok!
[Info]migration done, handle count: 2
```
### Step Migration
```shell
go run . dmc migrate --step=1
```
```shell
# output
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]migration start
[Info]migration:V20250214094800702CreateTableUsers
[Success]migration count: 1 version: V20250214094800702CreateTableUsers ok!
[Info]migration done, handle count: 1
```
### Rollback Migration
```shell
go run . dmc rollback --step=1
```
```shell
# output
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]rollback start
[Info]rollback(1):V20250214095503505AlterTableUsers
[Success]rollback count:1 version: V20250214095503505AlterTableUsers ok!
[Info]rollback done, handle count: 1
```
### Rollback All
```shell
go run . dmc rollback --all 
```
```shell
# output  
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]rollback start
[Info]rollback(99999999):V20250214095503505AlterTableUsers
[Success]rollback count:99999999 version: V20250214095503505AlterTableUsers ok!
[Info]rollback(99999999):V20250214094800702CreateTableUsers
[Success]rollback count:99999999 version: V20250214094800702CreateTableUsers ok!
[Info]rollback done, handle count: 2

```

### Repair Migration [New Feature]
**When the migration table cannot be migrated correctly due to operational errors, this command can be used**

If there are any issues with the status, it needs to be manually modified in the database
```shell
go run . dmc repair
```
```shell
[Info]check dc_migrations table
[Info]dc_migrations is ok
[Info]Fix unmatched items start  
[Success]Fix unmatched items count:0 
[Info]Fix missing items start
[Success]Fix missing items count:0
[Warning]If there are any issues with the status, it needs to be manually modified in the database
[Success]repair ok!

```
## Automatically generate migration description
### Automatically generate table structures for support based on databases
| id | tag                                      | already_migrated | created_at           | updated_at           | executed_at          | reverted_at          |
|----|------------------------------------------|------------------|----------------------|----------------------|----------------------|----------------------|
| 1  | v_2025_02_14_09_48_00_702_create_table_users | 1                | 2025-02-14 09:48:00.976 | 2025-02-14 10:05:20.698 | 2025-02-14 10:05:20.698 | 2025-02-14 10:04:50.403 |
| 2  | v_2025_02_14_09_55_03_505_alter_table_users | 0                | 2025-02-14 09:55:03.022 | 2025-02-14 10:04:50.392 | 2025-02-14 10:04:36.251 | 2025-02-14 10:04:50.392 |

### Automatically generate directory structure
```shell
$ tree example/
example/
├── go.mod
├── go.sum
└── dc_migrations
    ├── migration_v_2025_02_14_09_48_00_702_create_table_users.go
    ├── migration_v_2025_02_14_09_55_03_505_alter_table_users.go
    └── register.go
	└── dmc.go
    
$ cat example/dc_migrations/register.go 
package dc_migrations

import (
        "github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
    migrate.RegisterMigration("V20250214094800702CreateTableUsers", NewMigrateV20250214094800702CreateTableUsers())
        migrate.RegisterMigration("V20250214095503505AlterTableUsers", NewMigrateV20250214095503505AlterTableUsers())
}

```
### Automatically generate database List
```mysql
USE test;
SHOW TABLES;
```
| Tables in test       |
|----------------------|
| dc_migrations       |
| users                |