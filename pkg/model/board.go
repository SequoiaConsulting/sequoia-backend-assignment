package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

// Board declares database schema for boards table.
type Board struct {
	gorm.Model
	Name        string `gorm:"not null;size: 128;unique"`
	AdminUserID uint   `gorm:"not null;foreignkey;"`
	IsArchived  *bool  `gorm:"default: false;"`
}

// BeforeSave implements the GORM hook to perform input validation
func (board Board) BeforeSave() error {
	return validation.ValidateStruct(&board,
		validation.Field(&board.Name, validation.Required),
		validation.Field(&board.AdminUserID, validation.Required),
	)
}

// BeforeCreate implements the GORM to look for duplication entries
func (board Board) BeforeCreate(txn *gorm.DB) error {
	found := &Board{}
	err := found.FindByName(txn, board.Name)
	if err != nil {
		if err == ErrBoardNotFound {
			return nil
		}
		return err
	}
	return ErrDuplicateEntry
}

// FindByID tries to find a board with given ID. populates receiver object with board
// details if found, returns a non-nil error otherwise.
func (board *Board) FindByID(db *gorm.DB, id uint) error {
	err := db.Model(board).
		Where(&Board{Model: gorm.Model{ID: id}}).
		First(board).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrBoardNotFound
		}

		loggerFor("board").Warn().Err(err).Msg("unable to exec find by id query")
		return ErrInternalServerError
	}

	return nil
}

// FindByName tries to find a board with given name. populates receiver object with board
// details if found, returns a non-nil error otherwise.
func (board *Board) FindByName(db *gorm.DB, name string) error {
	err := db.Model(board).
		Where(&Board{Name: name}).
		First(board).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrBoardNotFound
		}

		loggerFor("board").Warn().Err(err).Msg("unable to exec find by name query")
		return ErrInternalServerError
	}

	return nil
}

// Create inserts the receiver object in the database
func (board *Board) Create(db *gorm.DB) error {
	err := db.Model(board).Create(board).Error
	if err != nil {
		if err == ErrDuplicateEntry {
			return err
		}

		loggerFor("board").Warn().Err(err).Msg("unable to exec create query")
		return ErrInternalServerError
	}

	return nil
}

// Update updates the receiver object in the database
func (board *Board) Update(db *gorm.DB) error {
	err := db.Model(board).
		Where(&Board{Model: gorm.Model{ID: board.ID}}).
		Update(board).
		Error

	if err != nil {
		loggerFor("board").Warn().Err(err).Msg("unable to exec create query")
		return ErrInternalServerError
	}

	return nil
}
