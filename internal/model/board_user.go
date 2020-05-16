package model

import (
	"github.com/jinzhu/gorm"
)

// BoardUser model to store boardID, userID
type BoardUser struct {
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	BoardID int `gorm:"not null;foreignkey"`
	UserID int `gorm:"not null;foreignkey"`
}

// POST /boarduser
// Create new board user
func (boardUser *BoardUser) Create(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(boardUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Add log
		return
	}
	db.Create(&boardUser)
	c.JSON(http.StatusOK, gin.H{"data": boardUser})

}

// GET /boardser/:id
// Find a boardUser
func (boardUser *BoardUser) FindByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(board).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": boardUser})
}

// UPDATE /boarduser/:id
// Update a boardUser
func (boardUser *BoardUser) Update(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).Update(boardUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			// Add log
		return	
	db.Model(&boardUser).Updates(boardUser)

	c.JSON(http.StatusOK, gin.H{"data": boardUser})
}

// DELETE /boarduser/:id
// Delete a boardUser
func (boardUser *BoardUser) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(boardUser).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
	db.Delete(&boardUser)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }


