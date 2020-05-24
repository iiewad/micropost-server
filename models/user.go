package models

import (
	"github.com/bxcodec/faker/v3"
	"github.com/iiewad/micropost-server/db"
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name  string
	Email string  `gorm:"type:varchar(100);unique_index"`
	UUID  *string `gorm:"primary_key;unique;not null"` // 设置会员号（member number）唯一并且不为空
}

// // BeforeCreate User
// func (user *UserModel) BeforeCreate(scope *gorm.Scope) error {
// 	uuid := uuid.New().String()
// 	scope.SetColumn("UUID", &uuid)
// 	return nil
// }

// UserSeed Seed
func UserSeed() {
	for i := 0; i < 100; i++ {
		uuid := faker.UUIDHyphenated()
		user := User{Name: faker.Name(), Email: faker.Email(), UUID: &uuid}
		db.DB.FirstOrCreate(&user, user)
	}
}
