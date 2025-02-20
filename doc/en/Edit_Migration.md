## Edit the migration file

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
1. **OpenEdit migration file on ide** 

   "./dc_migrations/migration_v_{Tag}_create_table_users.go"
1. **define the struct**
    ```go
    type Struct_Tag_CreateTableUsers struct{
        Id        int64 `gorm:"primaryKey;autoIncrement"`
        CreatedAt int64 `gorm:"autoCreateTime"`
        UpdatedAt int64 `gorm:"autoUpdateTime"`
        DeletedAt int64 `gorm:"index"`
    }
    ```
1. **write your migration code**
    ```go
    // Up is migration function
    func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().CreateTable(r.upStruct)
        if err != nil {
            return err
        }
        return nil
    }
    ```
1. **Write your rollback func**
    ```go
    // Down is rollback function
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx is the gorm.DB instance
        err := tx.Migrator().DropTable(r.downStruct)
        if err != nil {
            return err
        }
        return nil
    }
    ```
####  Alter column
1. **generate migration file**
    ```shell
    $ go run dmc.go gen --alter users
    ```
    ```shell
    # output
    [Info]check dc_migrations table
    [Info]dc_migrations is ok
    [Info]create migration start
    [Success]ok! file name :./dc_migrations/migration_v_2025_02_14_13_51_37_508_alter_table_users.go
    ```
1. **define the struct**
    ```go
    //The field data structure to be migrated and the table name correspondence will be automatically generated
    type StructV20250214135137508AlterTableUsersUp struct{
	    UserName        string `gorm:"gorm:"type:varchar(100);""`
	    NickName        string `gorm:"gorm:"type:varchar(120);""`
    }
   
   //To roll back a data structure, if a table or field is deleted, there is no need to declare the field. If a field is modified, it must be declared
   type StructV20250214135137508AlterTableUsersDown struct{
   }
    ```

1. **write your migration code**
 ```go
       // Up is migration function
       func (r *Migrate_Tag_CreateTableUsers) Up(tx *gorm.DB) error{
           //tx is the gorm.DB instance
          err := tx.Migrator().AddColumn(r.upStruct,"UserName")
          if err != nil {
            return err
          }
          err := tx.Migrator().AddColumn(r.upStruct,"NickName")
          if err != nil {
              return err
          }
          return nil
       }
```

1. **Write your rollback func**
```go
    // Down is rollback function
    func (r *Migrate_Tag_CreateTableUsers) Down(tx *gorm.DB) error{
        //tx is the gorm.DB instance
      err := tx.Migrator().DropColumn(r.downStruct, "UserName")
       if err != nil {
           return err
       }
   
       err = tx.Migrator().DropColumn(r.downStruct, "NickName")
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
