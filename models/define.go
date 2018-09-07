package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"gitlab.wallstcn.com/wscnbackend/ivankastd"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/toolkit"
)

var (
	db *gorm.DB
)

func InitModel(config ivankastd.ConfigMysql) {
	db = toolkit.CreateDB(config)
	db.AutoMigrate(
		&OperationLogs{},
	)
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func DB() *gorm.DB {
	return db
}

type OperationLogs struct {
	Id             int64 `grom:"primary_key"`
	MessageType    string
	Project        string
	Receiver       string
	SmsProject     string
	EmailTitle     string
	MessageContent string
	Status         string
	ErrorMsg       string
	CreatedAt      time.Time `gorm:"column:created_at;type:TIMESTAMP(6);default:CURRENT_TIMESTAMP(6);index"`
}
