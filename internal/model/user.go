package model

// User model to store id, name, email, password, isAdmin
type User struct {
	// gorm.Model
	ID           int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Name         string `gorm:"size:256;not null" json:"name"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"passwordHash"`
	IsAdmin      *bool  `gorm:"default: false" json:"isAdmin"`
}

func (*User) TableName() string {
	return "users"
}