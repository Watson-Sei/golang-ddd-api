package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../ddd.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Close() {
	db, err := db.DB()
	if err = db.Close(); err != nil {
		panic(err)
	}
}
