package models

import "github.com/iiewad/micropost-server/db"

// Init models
func Init() {
	db.DB.AutoMigrate(&User{})
	// UserSeed()
}
