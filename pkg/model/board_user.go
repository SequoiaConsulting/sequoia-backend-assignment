package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// BoardUser is a many to many assignment relationship between users
// and boards.
type BoardUser struct {
	BoardID uint `gorm:"primary_key;foreignkey;"`
	UserID  uint `gorm:"primary_key;foreignkey;"`
}

// BeforeSave implements the GORM hook to validate input before create/update.
func (bu BoardUser) BeforeSave() error {
	if bu.BoardID < 1 || bu.UserID < 1 {
		return errors.New("invalid board user mapping")
	}
	return nil
}

// BeforeCreate implements the GORM hook to check for duplicate entries before create query
func (bu BoardUser) BeforeCreate(txn *gorm.DB) error {
	err := txn.Model(&bu).Where(&bu).First(&bu).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}

		loggerFor("board_user").Warn().Err(err).Msg("unable to exec find query")
		return ErrInternalServerError
	}
	return ErrDuplicateEntry
}

// Create inserts the receiver object in the database
func (bu *BoardUser) Create(db *gorm.DB) error {
	err := db.Model(bu).Create(bu).Error
	if err != nil {
		if err == ErrDuplicateEntry {
			return err
		}

		loggerFor("board").Warn().Err(err).Msg("unable to exec create query")
		return ErrInternalServerError
	}

	return nil
}

// Exists checks if a user with given id is assigned to a board
func (bu *BoardUser) Exists(db *gorm.DB) error {
	err := db.Model(bu).Where(bu).First(bu).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrBoardUserNotFound
		}

		loggerFor("board_user").Warn().Err(err).Msg("unable to exec find board-user mapping query")
		return ErrInternalServerError
	}

	return nil
}

// Delete deletes the receiver object from the database
func (bu *BoardUser) Delete(db *gorm.DB) error {
	err := db.Unscoped().Model(bu).
		Where(&BoardUser{BoardID: bu.BoardID, UserID: bu.UserID}).
		Delete(bu).Error

	if err != nil {
		loggerFor("board").Warn().Err(err).Msg("unable to exec delete query")
		return ErrInternalServerError
	}

	return nil
}
