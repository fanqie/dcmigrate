这是一个基于 Gorm的数据库迁移工具

**[English](../../README.md) | 中文**
# Warning⚠️
- **目前版本还在实验中，请谨慎使用！！！**
# 快速开始
## 安装依赖
```shell
go get -u github.com/fanqie/dcmigrate
```
### 快速初始化迁移工具在你的项目
[操作指南文档](doc/zh_cn/Initialize_Guide.md)


### 目录结构
这个目录结构是初始化自动生成的
```shell
example/
├── dmc.go // 这个是dcmigrate 的命令行工具入口
├── dc_migrations // 迁移文件目录
  └── register.go // 这是dcmigration的迁移文件注册文件，由dcmigration自动生成和维护。请不要手动修改
  └── 20230301_000000_create_users_table.go // 这是dcmigration的迁移文件
  
├── go.mod
├── go.sum
└── ... you project files

```
### 如何链接数据库
打开“dmc.go”文件，修改数据库连接信息，然后运行dmc。去文件。您可以根据参考代码和Gorm官方文档配置数据库连接

**[Gorm连接数据库文档指南](https://gorm.io/docs/connecting_to_the_database.html)**
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

// Gorm连接数据库文档指南 doc:  https://gorm.io/docs/connecting_to_the_database.html

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
### 生成
#### 生成createTable迁移文件
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
#### 编辑迁移文件
[操作指南文档](doc/zh_cn/Edit_Migration.md)
#### 生成alterTable迁移文件
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
### 运行迁移
#### 全部迁移
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
#### 步进迁移
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
### 回滚迁移
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
### 回滚所有迁移
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
### 自动生成迁移描述
#### 基于数据库自动生成用于支持的表结构
| id | tag                                      | already_migrated | created_at           | updated_at           | executed_at          | reverted_at          |
|----|------------------------------------------|------------------|----------------------|----------------------|----------------------|----------------------|
| 1  | v_2025_02_14_09_48_00_702_create_table_users | 1                | 2025-02-14 09:48:00.976 | 2025-02-14 10:05:20.698 | 2025-02-14 10:05:20.698 | 2025-02-14 10:04:50.403 |
| 2  | v_2025_02_14_09_55_03_505_alter_table_users | 0                | 2025-02-14 09:55:03.022 | 2025-02-14 10:04:50.392 | 2025-02-14 10:04:36.251 | 2025-02-14 10:04:50.392 |

#### 自动生成目录结构
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
#### 自动生成数据库列表
```mysql
USE test;
SHOW TABLES;
```
| Tables in test       |
|----------------------|
| dc_migrations       |
| users                |