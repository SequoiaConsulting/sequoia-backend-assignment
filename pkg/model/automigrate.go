package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AutoMigrate migrates all schemas in database to the declared versions in
// the source code.
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		new(RevokedSession),
		new(User),
		new(Board),
		new(BoardUser),
		new(TaskStatus),
		new(Task),
	).Error

	if err != nil {
		return errors.Wrap(err, "unable to migrate database schemas")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.DefaultCost)
	IsAdmin := true
	// create a admin user by default
	err = db.Model(new(User)).FirstOrCreate(&User{
		Name:     "root",
		Email:    "root@admin.org",
		Password: string(hash),
		IsAdmin:  &IsAdmin,
	}).Error

	if err != nil {
		return errors.Wrap(err, "unable to create root user")
	}

	return nil
}
