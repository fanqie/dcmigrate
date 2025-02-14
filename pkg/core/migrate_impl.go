package core

import (
	"database/sql"
	"gorm.io/gorm"
)

type DcMigrateImpl interface {
	Up(tx *gorm.DB) error
	Down(tx *gorm.DB) error
	Register()
	UpAfter()
	DownAfter()
	GetTypeTag() string
	TableName() string
	GetData() *MigrateBasic
	SetRevertedAt(value sql.NullTime)
	GetRevertedAt() sql.NullTime
	GetExecutedAt() sql.NullTime
	SetExecutedAt(value sql.NullTime)
	GetAlreadyMigrated() bool
	SetAlreadyMigrated(value bool)
	SetTag(value string)
	GetTag() string
	SetId(value int)
	GetId() int
}
