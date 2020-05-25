package common

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

// DBInit *
func DBInit() {
	var err error
	DB, err = gorm.Open("mysql", "root:1005@/micropost-development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}

// DBClose DB
func DBClose() {
	defer DB.Close()
}
