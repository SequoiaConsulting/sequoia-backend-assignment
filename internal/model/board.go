package model

import (
	"github.com/jinzhu/gorm"
)

// Board model to store name, isArchived, userID
type Board struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	isArchived bool `gorm:"default:false;"`
	UserID int `gorm:"not null;foreignkey"`
}

// POST /board
// Create new board
func (board *Board) Create(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Add log
		return
	}
	db.Create(&board)
	c.JSON(http.StatusOK, gin.H{"data": board})

}

// GET /board/:id
// Find a board
func (board *Board) FindByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(board).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": board})
}

// UPDATE /board/:id
// Update a board
func (task *Task) Update(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).Update(board).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			// Add log
		return
	db.Model(&task).Updates(task)

	c.JSON(http.StatusOK, gin.H{"data": board})
}


// DELETE /board/:id
// Delete a board
func (board *Board) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(board).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
	db.Delete(&board)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }



