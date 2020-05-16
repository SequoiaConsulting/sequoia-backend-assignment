package model

import (
	"github.com/jinzhu/gorm"
)

// Status model to store id, status name, boardID
type Status struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	BoardID int `gorm:"not null;foreignkey"`
}  