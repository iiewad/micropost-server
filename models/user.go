package models

import (
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/iiewad/micropost-server/common"
	"github.com/iiewad/micropost-server/utils"
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name         string
	Email        string  `gorm:"type:varchar(100);unique_index"`
	UUID         *string `gorm:"primary_key;unique;not null"` // 设置会员号（member number）唯一并且不为空
	PasswordHash string  `gorm:"not null"`
	Microposts   []Micropost
}

// UserInfo Export
type UserInfo struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate User
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New().String()
	scope.SetColumn("UUID", &uuid)
	password, err := utils.PasswordSecret(user.PasswordHash)
	if err != nil {
		return err
	}
	scope.SetColumn("PasswordHash", password)
	return nil
}

// UserInfo Data
func (user *User) UserInfo() UserInfo {
	userInfo := &UserInfo{
		ID:        *user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return *userInfo

}

// AddUser Func
func AddUser(tx *gorm.DB, user User) (*string, error) {
	err := tx.Create(&user).Error
	userID := user.UUID
	return userID, err
}

// UserSeed Seed
func UserSeed() {
	common.DB.Unscoped().Delete(&User{})
	for i := 0; i < 100; i++ {
		uuid := faker.UUIDHyphenated()
		password := "password"
		passwordHash, err := utils.PasswordSecret(password)
		if err != nil {
			log.Fatal(err)
		}
		user := User{Name: faker.Name(), Email: faker.Email(), UUID: &uuid, PasswordHash: string(passwordHash)}
		common.DB.Create(&user)
	}
}
