package model

import (
	"time"
	
	"github.com/jinzhu/gorm"
	valid "github.com/go-ozzo/ozzo-validation/v4"
)

// Task model to store id, title, description, due date, assigned by, assigned to, statusID, boardID
type Task struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Title string `gorm:"size:256;not null"`
	Description string `gorm:"size:1024;not null"`
	DueDate *time.Time `gorm:"default:null"`
	CreatedBy string `gorm:"foreignkey:User"`
	AssignedBy string `gorm:"foreignkey:User"`
	AssignedTo string `gorm:"foreignkey:User"`
	StatusID int `gorm:"foreignkey:Status"`
	BoardID int `gorm:"foreignkey:Board"`
}

func (*Task) TableName() string {
	return "tasks"
}

func ValidateTask(task Task) error {
	return valid.ValidateStruct(
		valid.Field(&task.Title, valid.Required),
		valid.Field(&task.CreatedBy, valid.Required),
	)
}
