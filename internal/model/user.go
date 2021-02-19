package model

import (
	//valid "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"

)

// User model to store id, name, email, password, isAdmin
type User struct {
	gorm.Model
	// ID    int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Name  string `gorm:"size:256;not null" json:"name"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	IsAdmin *bool `gorm:"default: false" json:"isAdmin"`
	Token   string `gorm:"not null" json:"token"`
}

func (*User) TableName() string {
	return "users"
}

// func ValidateUser(user *User) error {
// 	return valid.ValidateStruct(
// 		valid.Field(&user.Name, valid.Required),
// 		valid.Field(&user.Email, valid.Required),
// 		valid.Field(&user.Password, valid.Required),
// 	)
// }
