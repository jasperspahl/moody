package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("moody.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("unable to connect to db: %v", err)
	}

	err = db.AutoMigrate(&Mood{}, &Entry{}, &Activity{})
	if err != nil {
		log.Panicf("failed to auto migrate: %v", err)
	}
}

func DB() *gorm.DB {
	return db
}
