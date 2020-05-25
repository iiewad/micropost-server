package common

import (
	"github.com/jinzhu/gorm"
)

// DB gorm
var DB *gorm.DB

// Init *
func Init() {
	DBInit()
}
