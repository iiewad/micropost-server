package models

import "github.com/iiewad/micropost-server/common"

// Result struct
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Init models
func Init() {
	common.DB.AutoMigrate(&User{})
	// UserSeed()
}
