package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task declares the table schema for a Task on any given board.
type Task struct {
	gorm.Model
	Title       string     `gorm:"size: 128;"`
	Description string     `gorm:"size: 1024;"`
	DueDate     *time.Time `gorm:"default: null;"`
	AssigneeID  *uint      `gorm:"foreignkey;"`
	AssignerID  *uint      `gorm:"foreignkey;"`
	BoardID     *uint      `gorm:"foreignkey;"`
	StatusID    *uint      `gorm:"foreignkey;"`

	Assignee User
	Assigner User
	Board    Board
	Status   TaskStatus
}
