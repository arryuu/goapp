package models

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func BindDb(bindDb *gorm.DB) {
	db = bindDb
	autoMigrate()
}

func autoMigrate() {
	err := db.AutoMigrate()
	if err != nil {
		return
	}
}
