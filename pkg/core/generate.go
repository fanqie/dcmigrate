package core

import (
	"fmt"
	"github.com/fanqie/dcmigrate/pkg/tpl"
	"github.com/fanqie/dcmigrate/pkg/utility"
	"gorm.io/gorm"
	"os"
	"strings"
)

type GenArgs struct {
	Action    string
	TableName string
}

func gen(genArgs GenArgs, migrationsManage *MigratesManage) (error, string) {

	mb := &MigrateBasic{}
	mb.genRecord(genArgs)
	var fileName string
	return Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(mb)
		if result.Error != nil {
			utility.ErrPrintf("insert error", result.Error)
			return result.Error
		}
		err := migrationsManage.RefreshMigrationsData(tx)
		if err != nil {
			return err
		}
		err, fileName = GenFile(mb, migrationsManage, genArgs)
		return err
	}), fileName
}
func GenFile(r *MigrateBasic, migrationsManage *MigratesManage, genArgs GenArgs) (error, string) {

	migrateFileName := fmt.Sprintf("./dc_migrations/migration_%s.go", r.Tag)
	registerFileName := "./dc_migrations/register.go"

	err := saveFile(migrateFileName, makeMigrateFile(r, genArgs))
	if err != nil {
		return err, ""
	}
	err = overwriteFile(registerFileName, refreshRenderRegisterCode(migrationsManage))
	if err != nil {
		return err, ""
	}
	return nil, migrateFileName
}
func refreshRenderRegisterCode(migrationsManage *MigratesManage) string {
	content := tpl.RegisterCode
	var registerMigration []string
	for _, migration := range migrationsManage.MigrateList {
		migrationCode := fmt.Sprintf("migrate.RegisterMigration(\"%s\", NewMigrate%s())", migration.GetTypeTag(), migration.GetTypeTag())
		registerMigration = append(registerMigration, migrationCode)
	}
	return strings.Replace(content, "{{RegisterMigration}}", strings.Join(registerMigration, "\n\t"), -1)
}
func makeMigrateFile(r *MigrateBasic, genArgs GenArgs) string {
	var content string
	switch genArgs.Action {
	case "create":
		content = tpl.MigrationCreateTableCode
		break
	case "alter":
		content = tpl.MigrationAlterTableCode
		break
	default:
		panic("not support action")
	}
	content = strings.Replace(content, "{{Tag}}", r.Tag, -1)
	content = strings.Replace(content, "{{TypeTag}}", r.GetTypeTag(), -1)
	content = strings.Replace(content, "{{TableName}}", genArgs.TableName, -1)
	return content
}

// saveFile saves the content to the specified file path.
func saveFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		utility.ErrPrintf("Error creating file:%v", err)
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		utility.ErrPrintf("Error writing to file:%v", err)
		return err
	}
	return err
	//defer file.Close().Error()
}
func overwriteFile(path string, content string) error {
	err := os.Remove(path)
	if err != nil {
		utility.ErrPrintf("Error deleting file:%v", err)
		return err
	}
	return saveFile(path, content)

}
