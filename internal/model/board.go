package model

import (
	"github.com/jinzhu/gorm"
)

// Board model to store name, isArchived, userID
type Board struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	isArchived bool `gorm:"default:false;"`
	UserID int `gorm:"not null;foreignkey"`
}