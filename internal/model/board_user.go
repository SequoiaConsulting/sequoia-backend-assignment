package model

import (
	"github.com/jinzhu/gorm"
)

// BoardUser model to store boardID, userID
type BoardUser struct {
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	BoardID int `gorm:"not null;foreignkey"`
	UserID int `gorm:"not null;foreignkey"`
}