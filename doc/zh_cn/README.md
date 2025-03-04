**[English](../../README.md) | 中文**

dcmigrate是一个基于gorm的数据库迁移工具， 与gorm migrate机制为互补关系， 并无功能冲突，旨在提供一个更友好数据库迁移工具
### 特性

- 支持mysql、sqlite、postgres、sqlserver、tidb、clickhouse等数据库
- 命令行工具
- 依赖gorm
- 支持按步长执行迁移文件
- 支持回滚功能
- 可查看迁移列表以及状态
- 显式的“up”、“down” 函数实现
- 支持修复迁移记录

# Tooltip ⚠

- **目前版本较稳定，还会继续更新，建议保持到版本更新**

# 快速开始

## 安装依赖

```shell
go get -u github.com/fanqie/dcmigrate
```

## 快速初始化迁移工具在你的项目

[操作指南文档](Initialize_Guide.md)

## 目录结构

这个目录结构是初始化自动生成的

```shell
example/

├── dc_migrations // 迁移文件目录
  └── register.go // 这是dcmigration的迁移文件注册文件，由dcmigration自动生成和维护。请不要手动修改
  └── 20230301_000000_create_users_table.go // 这是dcmigration的迁移文件
  └── dmc.go // 这个是dcmigrate 的命令行工具入口
├── go.mod
├── go.sum
└── ... you project files

```
## 命令行使用说明
**您必须完成初始化操作**
```shell
go run . dmc  --help
```
```shell
[Info]check dc_migrations table
[Info]dc_migrations is ok
用法:
   [命令]

可用命令:
  completion  为指定 Shell 生成自动补全脚本
  gen         生成一个新的核心文件
        语法：dmc gen  [--create|--alter]  {表名}
        示例：`dmc gen --create user` // 或者 `dmc gen --alter user`
  help        显示任意命令的帮助信息
  list        显示所有迁移记录
  migrate     迁移所有新的迁移文件版本或目标步数版本
  rollback    回滚历史迁移
        语法：dmc rollback  [--step=1] [--all]  {表名}
        示例：`dmc rollback --step=2` // 或者 `dmc rollback --all`

选项:
  -h, --help   显示此帮助信息

使用 " [命令] --help" 查看有关某个命令的更多信息。

```

## 如何链接数据库

打开“dc_migrations/dmc.go”文件，修改数据库连接信息，然后运行dmc。去文件。您可以根据参考代码和Gorm官方文档配置数据库连接

**[Gorm连接数据库文档指南](https://gorm.io/docs/connecting_to_the_database.html)**

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
## 在你的项目中使用dcmigrate

```shell
package main
import "dc_migrations"
func main() {
	if dc_migrations.TryStartUpDcMigrate() {
		return
	}
	//todo: 你的代码在这里...
	// 如：gin.Run()
}
```
## 生成

### 生成createTable迁移文件

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

### 编辑迁移文件

[操作指南文档](Edit_Migration.md)

### 生成alterTable迁移文件

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

## 显示迁移列表

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

## 运行迁移

### 全部迁移

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

### 步进执行迁移

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

### 回滚迁移

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

### 回滚所有迁移

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
### 修复迁移记录 [新特性]


**当迁移表由于操作错误而无法正确执行迁移时，可以使用此命令**

如果状态有任何问题，则需要在数据库中手动修改

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

## 自动生成迁移描述

### 基于数据库自动生成用于支持的表结构

| id | tag                                          | already_migrated | created_at              | updated_at              | executed_at             | reverted_at             |
|----|----------------------------------------------|------------------|-------------------------|-------------------------|-------------------------|-------------------------|
| 1  | v_2025_02_14_09_48_00_702_create_table_users | 1                | 2025-02-14 09:48:00.976 | 2025-02-14 10:05:20.698 | 2025-02-14 10:05:20.698 | 2025-02-14 10:04:50.403 |
| 2  | v_2025_02_14_09_55_03_505_alter_table_users  | 0                | 2025-02-14 09:55:03.022 | 2025-02-14 10:04:50.392 | 2025-02-14 10:04:36.251 | 2025-02-14 10:04:50.392 |

### 自动生成目录结构

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

### 自动生成数据库列表

```mysql
USE test;
SHOW TABLES;
```

| Tables in test |
|----------------|
| dc_migrations  |
| users          |