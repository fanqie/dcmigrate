package main

import (
	"os"
	"path/filepath"

	"github.com/fanqie/dcmigrate/pkg/utility"
)

func main() {
	dir := "dc_migrations"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			utility.ErrPrintf("folder is fail: %s,please retry", dir)
			return
		}
		utility.SuccessPrintf("create folder success: %s", dir)
	} else {
		utility.ErrPrintf("folder is exists: %s,please remove it", dir)
	}

	registerPath := filepath.Join("dc_migrations", "register.go")
	if err := checkAndCreateFile(registerPath, registerContent()); err != nil {
		utility.ErrPrintf("create register.go fail:", err)
	}
	dmcPath := filepath.Join("dc_migrations", "dmc.go")
	if err := checkAndCreateFile(dmcPath, dmcContent()); err != nil {
		utility.ErrPrintf("create dmc.go fail:", err)
	}

}

func checkAndCreateFile(path, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
		utility.SuccessPrintf("create file success: %s", path)
	} else {
		utility.ErrPrintf("file is exists: %s,please remove it", path)
	}
	return nil
}

func dmcContent() string {
	return `package dc_migrations

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
			//Logger: logger.Default.LogMode(logger.Info),
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


`
}

func registerContent() string {
	return `package dc_migrations

import (
	"github.com/fanqie/dcmigrate/pkg"
)

func Register(migrate *pkg.DcMigrate) {
}
`
}
