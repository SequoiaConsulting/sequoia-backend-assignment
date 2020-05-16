package model

import (
	"github.com/jinzhu/gorm"
)

// User model to store id, name, email, password, isAdmin
type User struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin *bool `gorm:"default: false"`
}