package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Task declares the table schema for a Task on any given board.
type Task struct {
	gorm.Model
	Title       string     `gorm:"size: 128; not null"`
	Description string     `gorm:"size: 1024;"`
	BoardID     uint       `gorm:"foreignkey;"`
	OwnerID     uint       `gorm:"not null;foreignkey;"`
	DueDate     *time.Time `gorm:"default: null;"`
	AssigneeID  *uint      `gorm:"foreignkey;"`
	AssignerID  *uint      `gorm:"foreignkey;"`
	StatusID    *uint      `gorm:"foreignkey;"`

	Owner    User
	Assignee User
	Assigner User
	Board    Board
	Status   TaskStatus
}

// BeforeUpdate implements the GORM hook to validate input before updating database
func (task Task) BeforeUpdate() error {
	if task.BoardID < 1 || task.OwnerID < 1 {
		errors.New("we don't accept orphaned tasks")
	}

	return validation.ValidateStruct(&task,
		validation.Field(&task.Title, validation.Required, validation.Length(1, 128)),
	)
}

// Create inserts the task in database
func (task *Task) Create(db *gorm.DB) error {
	err := db.Model(task).Create(task).Error
	if err != nil {
		loggerFor("task").Warn().Err(err).Msg("unable to exec create query")
		return ErrInternalServerError
	}

	return nil
}
