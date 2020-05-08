package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// RevokedSession stores the revoked JWT token along with their expiry.
// This will be used by the authentication module for denying access to
// revoked tokens.
type RevokedSession struct {
	Token   string    `gorm:"primary_key;"`
	Expires time.Time `gorm:"not null;"`
}

// BeforeSave implements the GORM hook for validating data before saving
// to the database.
func (session RevokedSession) BeforeSave() error {
	if session.Token == "" || time.Since(session.Expires) > 0 {
		return errors.New("invalid data passed to RevokedSession")
	}
	return nil
}

// BeforeCreate implements the GORM hook to check if entry is already present
// in the database.
func (session RevokedSession) BeforeCreate(txn *gorm.DB) error {
	found := &RevokedSession{}
	err := found.FindByToken(txn, session.Token)
	if err != nil {
		if err == ErrInternalServerError {
			return err
		}
		return nil
	}
	return ErrDuplicateEntry
}

// FindByToken finds the revoked session instance using its token
func (session *RevokedSession) FindByToken(db *gorm.DB, token string) error {
	err := db.Model(session).Where(&RevokedSession{Token: session.Token}).First(session).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("session not found")
		}

		loggerFor("revoked_session").Warn().Err(err).Msg("unable to exec find by token query")
		return ErrInternalServerError
	}

	return nil
}

// Create inserts the given record into the database table.
func (session *RevokedSession) Create(db *gorm.DB) error {
	err := db.Model(session).Create(session).Error
	if err != nil {
		if err == ErrDuplicateEntry {
			return err
		}
		loggerFor("revoked_session").Warn().Err(err).Msg("unable to exec insert query")
		return ErrInternalServerError
	}
	return nil
}
