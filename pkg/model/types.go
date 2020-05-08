package model

import "github.com/pkg/errors"

var (
	// ErrDuplicateEntry indicates that query tried to create (insert) a duplicate
	// entry into the database.
	ErrDuplicateEntry = errors.New("duplicate entry")
	// ErrInternalServerError is returned when an unexpected non-recoverable error
	// is encountered.
	ErrInternalServerError = errors.New("internal server error")
	// ErrUserNotFound is returned when user with given filters could not be found.
	ErrUserNotFound = errors.New("user not found")
	// ErrSessionNotFound is returned when session with given token could not be found.
	ErrSessionNotFound = errors.New("session not found")
)
