package model

import (
	"github.com/jinzhu/gorm"
	// valid "github.com/go-ozzo/ozzo-validation/v4"
)

// Status model to store id, status name, boardID
type Status struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	BoardID int `gorm:"not null;foreignkey:Board"`
}  

func (*Status) TableName() string {
	return "statuses"
}

// func ValidateStatus(status Status) error {
// 	return valid.ValidateStruct(
// 		&status,
// 		valid.Field(&status.Name, valid.Required),
// 		valid.Field(&status.BoardID, valid.Required),
// 	)
// }
