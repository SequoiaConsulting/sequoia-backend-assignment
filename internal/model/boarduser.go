package model

import (
	"github.com/jinzhu/gorm"
	valid "github.com/go-ozzo/ozzo-validation/v4"
)

// BoardUser model to store boardID, userID
// It has many-to-many assignment
type BoardUser struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	BoardID int `gorm:"not null;foreignkey:Board"`
	UserID int `gorm:"not null;foreignkey:User"`
}

func (*BoardUser) TableName() string {
	return "boardusers"
}

func ValidateBoardUser(boardUser BoardUser) error {
	return valid.ValidateStruct(
		&boardUser,
		valid.Field(&boardUser.BoardID, valid.Required),
		valid.Field(&boardUser.UserID, valid.Required),
	)
}