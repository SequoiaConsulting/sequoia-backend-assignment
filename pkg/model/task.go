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

// Update updates the task in database
func (task *Task) Update(db *gorm.DB) error {
	err := db.Model(task).Where(&Task{Model: gorm.Model{ID: task.ID}}).Update(task).Error
	if err != nil {
		loggerFor("task").Warn().Err(err).Msg("unable to exec update query")
		return ErrInternalServerError
	}

	return nil
}

// FindByID finds a task by its ID
func (task *Task) FindByID(db *gorm.DB) error {
	err := db.Model(task).Where(&Task{Model: gorm.Model{ID: task.ID}}).First(task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrTaskNotFound
		}

		loggerFor("task").Warn().Err(err).Msg("unable to exec find by id query")
		return ErrInternalServerError
	}

	return nil
}

// Delete deletes the task in database
func (task *Task) Delete(db *gorm.DB) error {
	err := db.Model(task).Where(&Task{Model: gorm.Model{ID: task.ID}}).Delete(task).Error
	if err != nil {
		loggerFor("task").Warn().Err(err).Msg("unable to exec delete query")
		return ErrInternalServerError
	}

	return nil
}
