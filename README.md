this is a database migration tool based on Gorm
# Warning⚠️
- **This version is still in the experimental stage, please choose carefully！！！**
# Quick start
## Install dependencies
```shell
go get -u github.com/fanqie/dcmigrate
```
### Quickly Initialize Your Project
[Guide Doc](doc/en/Initialize_Guide.md)


### Directory Structure
This is the directory structure you obtained after initializing the project
```shell
example/
├── dmc.go // This is the command-line tool for gormMigration
├── dc_migrations // This is the migration file directory for gormMigration
  └── register.go // This is the migration file registration file for gormMigration, which is automatically generated and maintained by gormMigration. Please do not manually modify it
  └── 20230301_000000_create_users_table.go // This is the migration file for gormMigration
  
├── go.mod
├── go.sum
└── ... you project files

```
### Connect to Database
Open the "dmc. go" file, modify the database connection information, and then run the dmc. go file. You can configure the database connection according to the reference code and the official Gorm documentation

**[Gorm Connecting Database Doc Guide](https://gorm.io/docs/connecting_to_the_database.html)**
```shell
package main

import (
	"github.com/fanqie/dcmigrate-example/dc_migrations"
	"github.com/fanqie/dcmigrate/pkg"
	"github.com/fanqie/dcmigrate/pkg/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dcMigrate := pkg.NewDcMigrate(true)
	dc_migrations.Register(dcMigrate)
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
### Generate
#### Generate a createTable migration file
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
#### Edit the migration file
[Guide Doc](doc/en/Edit_Migration.md)
#### Generate a alterTable migration file
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
### Run Migration
#### Migrate All
```shell
$ go run dmc.go migrate       
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
#### Step Migration
```shell
$ go run dmc.go migrate --step=1
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
$ go run dmc.go rollback --step=1
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
$ go run dmc.go rollback --all 
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
### Automatically generate product description
####  Automatically generated support table structure on the basis of the database
| id | tag                                      | already_migrated | created_at           | updated_at           | executed_at          | reverted_at          |
|----|------------------------------------------|------------------|----------------------|----------------------|----------------------|----------------------|
| 1  | v_2025_02_14_09_48_00_702_create_table_users | 1                | 2025-02-14 09:48:00.976 | 2025-02-14 10:05:20.698 | 2025-02-14 10:05:20.698 | 2025-02-14 10:04:50.403 |
| 2  | v_2025_02_14_09_55_03_505_alter_table_users | 0                | 2025-02-14 09:55:03.022 | 2025-02-14 10:04:50.392 | 2025-02-14 10:04:36.251 | 2025-02-14 10:04:50.392 |

#### Automatically generate directory structure
```shell
$ tree example/
example/
├── dmc.go
├── go.mod
├── go.sum
└── dc_migrations
    ├── migration_v_2025_02_14_09_48_00_702_create_table_users.go
    ├── migration_v_2025_02_14_09_55_03_505_alter_table_users.go
    └── register.go
    
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
#### Automatically generate database List
```mysql
USE test;
SHOW TABLES;
```
| Tables in test       |
|----------------------|
| dc_migrations       |
| users                |