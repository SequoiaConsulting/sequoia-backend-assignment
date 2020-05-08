package model

import (
	"github.com/jinzhu/gorm"
)

// Board declares database schema for boards table.
type Board struct {
	gorm.Model
	Name string `gorm:"not null;size: 128;"`
}
