package core

import (
	"fmt"
	"github.com/fanqie/dcmigrate/pkg/utility"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

func DefinedCommand(migrationsManage *MigratesManage, migrations map[string]DcMigrateImpl) {
	rootCmd := &cobra.Command{}
	genCommand := &cobra.Command{
		Use: "gen",
		Short: "generate a new core file" +
			"\n\tsyntax：dmc gen  [--create|--alter]  {table_name}" +
			"\n\tusage：`dmc gen --create user` //or `dmc gen --alter user`",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utility.ErrPrint("syntax error, dmc gen [--create|--alter]  {table_name}")
				return
			}
			action := ""
			if cmd.Flags().Changed("create") {
				action = "create"
			} else if cmd.Flags().Changed("alter") {
				action = "alter"
			}
			if action == "" {
				utility.ErrPrint("syntax error, dmc gen [--create|--alter] {table_name}")
			}
			tableName := args[0]

			if tableName == "" {
				utility.ErrPrint("tableName is required")
				return
			}
			utility.InfoPrint("create migration start")
			err, fileName := gen(GenArgs{
				Action:    action,
				TableName: tableName,
			}, migrationsManage)
			if err != nil {
				utility.ErrPrint(err.Error())
				return
			}
			utility.SuccessPrintf("ok! file name :%s", fileName)
		},
	}
	genCommand.Flags().Bool("create", false, "set action to create the table,is default")
	genCommand.Flags().Bool("alter", false, "set action to alter the table")
	rootCmd.AddCommand(genCommand)

	rollbackCommand := &cobra.Command{
		Use: "rollback",
		Short: "rollback history migrates" +
			"\n\tsyntax：dmc rollback  [--step=1] [--all]  {table_name}" +
			"\n\tusage：`dmc rollback --step=2` //or `dmc rollback --all`",
		Run: func(cmd *cobra.Command, args []string) {
			var step int
			if cmd.Flags().Changed("step") {
				value, _ := cmd.Flags().GetInt("step")
				step = value
			} else {
				step = 1
				if cmd.Flags().Changed("all") {
					isAll, _ := cmd.Flags().GetBool("all")
					if isAll {
						step = 99999999
					} else {
						utility.InfoPrint("step is required, use default value 1")
						step = 1
					}
				}

			}

			err := RollbackHandle(step, *migrationsManage, migrations)
			if err != nil {
				utility.ErrPrintf("migrate error, %v", err.Error())
				return
			}
		},
	}
	rollbackCommand.Flags().IntP("step", "s", 1, "The default is to rollback one version. You can specify the number of versions to be rolled back and execute them in reverse order!")
	rollbackCommand.Flags().BoolP("all", "a", false, "rollback all migrations")
	// 添加子命令
	rootCmd.AddCommand(rollbackCommand)

	migrateCommand := &cobra.Command{
		Use:   "migrate",
		Short: "all new migration file versions will be migrated or target step size version",
		Run: func(cmd *cobra.Command, args []string) {
			var step int
			if cmd.Flags().Changed("step") {
				value, err := cmd.Flags().GetInt("step")
				if err != nil {
					utility.ErrPrintf("step is required, %v", err.Error())
					return
				}
				step = value
			}
			err := MigrateHandle(step, *migrationsManage, migrations)
			if err != nil {
				utility.ErrPrintf("migrate error, %v", err.Error())
				return
			}
		},
	}

	migrateCommand.Flags().IntP("step", "s", 1, "By default, all new migration file versions will be migrated, and you can also set the migration step size")
	rootCmd.AddCommand(migrateCommand)

	listCommand := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "status"},
		Short:   "show all migrations record",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrationsManage.RefreshMigrationsData(Db)
			if err != nil {
				return
			}
			//migrationsManage.MigrateList
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)

			// 设置表头和内容
			t.AppendHeader(table.Row{"ID", "Tag", "CreatedAt", "AlreadyMigrated"})
			var rows []table.Row
			for _, migtation := range migrationsManage.MigrateList {

				createdAt := migtation.CreatedAt.Format("2006-01-02 15:04:05")
				status := color.New(color.FgHiBlue).Sprint("Pending")
				if migtation.AlreadyMigrated {
					status = color.New(color.FgHiGreen).Sprint("☑ Yes!")
				}
				rows = append(rows, table.Row{
					migtation.ID,
					migtation.Tag,
					createdAt,
					status,
				})
			}
			t.AppendRows(rows)

			// 自定义样式
			t.SetStyle(table.StyleLight) // 内置简约样式

			t.Render() // 输出表格
		},
	}
	rootCmd.AddCommand(listCommand)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
