## 如何编写迁移文件使用说明

###  如何编写迁移文件
####  Create table
1. 生成一个迁移创建表的文件
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
1. **在你的ide中打开生成好的迁移文件** 

   "./dc_migrations/migration_v_{Tag}_create_table_users.go"
1. **声明数据结构**
    ```go
    type Struct_Tag_CreateTableUsers struct{
        Id        int64 `gorm:"primaryKey;autoIncrement"`
        CreatedAt int64 `gorm:"autoCreateTime"`
        UpdatedAt int64 `gorm:"autoUpdateTime"`
        DeletedAt int64 `gorm:"index"`
    }
    ```
1. **写入你的迁移代码**
    ```go
    // Up 是迁移处理函数
    func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
        //tx 是 gorm.DB 实例
        err := tx.Migrator().CreateTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
1. **Write your rollback func**
    ```go
    // Down 是回滚处理函数
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx 是 gorm.DB 实例
        err := tx.Migrator().DropTable(r.currentTable)
        if err != nil {
            return err
        }
        return nil
    }
    ```
####  修改活表字段
1. **生成一个修改迁移文件**
    ```shell
    $ go run dmc.go gen --alter users
    ```
    ```shell
    # 输出
    [Info]check dc_migrations table
    [Info]dc_migrations is ok
    [Info]create migration start
    [Success]ok! file name :./dc_migrations/migration_v_2025_02_14_13_51_37_508_alter_table_users.go
    ```
1. **声明数据结构**
    ```go
    type StructV20250214135137508AlterTableUsers struct{
	    UserName        string `gorm:"gorm:"type:varchar(100);""`
	    NickName        string `gorm:"gorm:"type:varchar(120);""`
    }
    ```

1. **编写你的迁移代码**
 ```go
       // Up 迁移处理函数
       func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
//tx 是 gorm.DB 实例
          err := tx.Migrator().AddColumn(r.currentTable,"UserName")
          if err != nil {
            return err
          }
          err := tx.Migrator().AddColumn(r.currentTable,"NickName")
          if err != nil {
              return err
          }
          return nil
       }
```

1. **Write your rollback func**
```go
    // Down 是回滚处理函数
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx 是 gorm.DB 实例
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
```
###  如何修改表名？
直接修改返回值即可
```go
func (*Struct_Tag_CreateTableUsers) TableName() string {
	return "users" //你可以改为你的目标数据库表的名称
}
```

###  [Gorm 字段标签规则](https://gorm.io/docs/models.html#Fields-Tags)
