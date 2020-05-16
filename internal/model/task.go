package model

import (
	"github.com/jinzhu/gorm"
)

// Task model to store id, title, description, due date, assigned by, assigned to, statusID, boardID
type Task struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` //why?
	Title string `gorm:"size:256;not null"`
	Description string `gorm:"size:1024;not null"`
	DueDate *time.Time `gorm:"default:null"`
	AssignedBy string `gorm:"foreignkey"`
	AssignedTo string `gorm:"foreignkey"`
	StatusID int `gorm:"foreignkey"`
	BoardID int `gorm:"foreignkey"`
}