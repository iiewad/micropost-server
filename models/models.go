package models

import "github.com/iiewad/micropost-server/common"

// Init models
func Init() {
	common.DB.AutoMigrate(&User{})
	// UserSeed()
}
