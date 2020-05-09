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
