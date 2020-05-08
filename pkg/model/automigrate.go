package model

import "github.com/jinzhu/gorm"

// AutoMigrate migrates all schemas in database to the declared versions in
// the source code.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(User),
		new(Board),
		new(TaskStatus),
		new(Task),
	).Error
}
