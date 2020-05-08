package model

import "time"

// RevokedSession stores the revoked JWT token along with their expiry.
// This will be used by the authentication module for denying access to
// revoked tokens.
type RevokedSession struct {
	Token   string    `gorm:"primary_key;"`
	Expires time.Time `gorm:"not null;"`
}
