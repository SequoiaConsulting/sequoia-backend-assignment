package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

// TaskStatus declares table schema to store task status types for
// any specific board.
type TaskStatus struct {
	gorm.Model
	BoardID uint   `gorm:"unique_index:idx_status;foreignkey;not null"`
	Title   string `gorm:"unique_index:idx_status;size:64;not null"`

	Board Board
}

// BeforeUpdate implements the GORM hook to validate input before applying to database
func (status TaskStatus) BeforeUpdate() error {
	return validation.ValidateStruct(&status,
		validation.Field(&status.BoardID, validation.Required),
		validation.Field(&status.Title, validation.Required),
	)
}

// BeforeCreate checks for duplicate entries
func (status TaskStatus) BeforeCreate(txn *gorm.DB) error {
	err := status.Find(txn)
	if err != nil {
		if err == ErrTaskStatusNotFound {
			return nil
		}
		return ErrInternalServerError
	}

	return ErrDuplicateEntry
}

// Find finds the task status with given title and board_id
func (status *TaskStatus) Find(db *gorm.DB) error {
	err := db.Model(status).Where(status).Find(status).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrTaskStatusNotFound
		}
		loggerFor("task_status").Warn().Err(err).Msg("unable to exec find query")
		return ErrInternalServerError
	}

	return nil
}

// Create inserts the receiver object in the database
func (status *TaskStatus) Create(db *gorm.DB) error {
	err := db.Model(status).Create(status).Error
	if err != nil {
		if err == ErrDuplicateEntry {
			return err
		}

		loggerFor("task_status").Warn().Err(err).Msg("unable to exec create query")
		return ErrInternalServerError
	}

	return nil
}

// Delete deletes the receiver object in the database
func (status *TaskStatus) Delete(db *gorm.DB) error {
	err := db.Model(status).Where(&TaskStatus{Model: status.Model}).Delete(status).Error
	if err != nil {
		loggerFor("task_status").Warn().Err(err).Msg("unable to exec delete query")
		return ErrInternalServerError
	}

	return nil
}
