package model

import "github.com/jinzhu/gorm"

// User declares the model for the user table in database
type User struct {
	gorm.Model
	Name     string `gorm:"size:128; not null;"`
	Email    string `gorm:"size:256; not null; unique;"`
	Password string `gorm:"size:60; not null"`
	Admin    bool
}
