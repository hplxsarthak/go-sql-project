package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This whole file will return a db which is a connection with database

var (
	db *gorm.DB
)

func Connect() {
	d,err := gorm.Open("mysql", "root:Sarthak@123@/medhplx?charset=utf8mb4&parseTime=True")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}