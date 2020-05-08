package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jinzhu/gorm"
)

// User declares the model for the user table in database
type User struct {
	gorm.Model
	Name     string `gorm:"size:128; not null;"`
	Email    string `gorm:"size:256; not null;"`
	Password string `gorm:"size:60; not null;"`
	IsAdmin  *bool  `gorm:"default: false;"`
}

// BeforeSave implements the GORM hook for validating user entries before inserting
// or updating them in the database.
func (user User) BeforeSave() error {
	return user.Validate()
}

// BeforeCreate implements the GORM hook for checking that user doesn't already exists
// If already exists, it returns ErrDuplicateEntry
func (user User) BeforeCreate(txn *gorm.DB) error {
	found := &User{}
	err := found.FindByEmail(txn, user.Email)
	if err == nil {
		return ErrDuplicateEntry
	}

	if err == ErrUserNotFound {
		return nil
	}

	return err
}

// Validate validates input fields in the user struct and returns the error if any
func (user *User) Validate() error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Name, validation.Required, validation.Length(1, 128), is.Alpha),
		validation.Field(&user.Email, validation.Required, validation.Length(5, 256), is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(60, 60)),
	)
}

// Create inserts the reciever object in the database table.
func (user *User) Create(db *gorm.DB) error {
	err := db.Model(user).Create(user).Error
	if err != nil {
		if err == ErrDuplicateEntry {
			return err
		}
		loggerFor("user").Warn().Err(err).Msg("unable to exec insert query")
		return ErrInternalServerError
	}
	return nil
}

// FindByEmail looks up the user by email address in the database. If found, it populates
// the receiver instance. Returns an other otherise.
func (user *User) FindByEmail(db *gorm.DB, email string) error {
	err := db.Model(user).Where(&User{Email: email}).First(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrUserNotFound
		}
		loggerFor("user").Warn().Err(err).Msg("unable to exec find by email query")
		return ErrInternalServerError
	}
	return nil
}

// FindByID looks up the user by its ID in the database. If found, it populates
// the receiver instance. Returns an other otherise.
func (user *User) FindByID(db *gorm.DB, id uint) error {
	err := db.Model(user).
		Where(&User{Model: gorm.Model{ID: id}}).
		First(user).
		Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrUserNotFound
		}
		loggerFor("user").Warn().Err(err).Msg("unable to exec find by ID query")
		return ErrInternalServerError
	}
	return nil
}

// Update updates the row in the table with values from receiver object. It takes the
// primary key from receiver object to lookup the row.
func (user *User) Update(db *gorm.DB) error {
	err := db.Model(user).
		Where(&User{Model: gorm.Model{ID: user.ID}}).
		Update(user).
		Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrUserNotFound
		}
		loggerFor("user").Warn().Err(err).Msg("unable to exec update query")
		return ErrInternalServerError
	}
	return nil
}
