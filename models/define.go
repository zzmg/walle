package models

import (
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
		&InServiceUser{},
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

type InServiceUser struct {
	Id        int64  `gorm:"primary_key"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
}
