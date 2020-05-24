package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

// DB GORM
var DB *gorm.DB

// Init DB
func Init() {
	var err error
	DB, err = gorm.Open("mysql", "root:1005@/micropost-development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}

// Close DB
func Close() {
	defer DB.Close()
}
