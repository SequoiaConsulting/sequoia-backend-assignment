package model

import (
	"github.com/jinzhu/gorm"
	// valid "github.com/go-ozzo/ozzo-validation/v4"
)

// Board model to store name, isArchived, userID
type Board struct {
	gorm.Model
	//ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	isArchived bool `gorm:"default:false;"`
	UserID int `gorm:"not null;foreignkey:User"`
}

func (*Board) TableName() string {
	return "boards"
}


// func ValidateBoard(board Board) error {
// 	return valid.ValidateStruct(
// 		&board,
// 		valid.Field(&board.Name, valid.Required),
// 		valid.Field(&board.isArchived, valid.Required),
// 		valid.Field(&board.UserID, valid.Required),
// 	)
// }