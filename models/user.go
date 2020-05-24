package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iiewad/micropost-server/db"
	"github.com/jinzhu/gorm"
)

// UserModel schema
type UserModel struct {
	gorm.Model
	Name         string
	Email        string  `gorm:"type:varchar(100);unique_index"`
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
}

// BeforeCreate User
func (user *UserModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("MemberNumber", uuid.New())
	return nil
}

// UserSeed Seed
func UserSeed() {
	user := UserModel{Name: "张三", Email: "zhangsan@test.local"}
	db.DB.FirstOrCreate(&user, user)
	fmt.Println(&user)
}
