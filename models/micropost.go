package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Micropost struct
type Micropost struct {
	gorm.Model
	UUID    *string `gorm:"primary_key;unique;not null"`
	UserID  *string `gorm:"not null"`
	Content *string `gorm:"type:text;not null;size:255"`
}

// BeforeCreate Micropost
func (user *Micropost) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New().String()
	scope.SetColumn("UUID", &uuid)
	return nil
}

// AddMicropost *
func AddMicropost(tx *gorm.DB, micropost Micropost) (*string, error) {
	err := tx.Create(&micropost).Error
	micropostID := micropost.UUID
	return micropostID, err
}
