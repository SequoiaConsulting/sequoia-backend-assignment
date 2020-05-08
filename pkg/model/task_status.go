package model

// TaskStatus declares table schema to store task status types for
// any specific board.
type TaskStatus struct {
	BoardID uint   `gorm:"primary_key; foreignkey;"`
	Title   string `gorm:"primary_key; size:64;"`

	Board Board
}
