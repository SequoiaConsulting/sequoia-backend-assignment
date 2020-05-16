package model

import (
	"github.com/jinzhu/gorm"
)

// Status model to store id, status name, boardID
type Status struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	BoardID int `gorm:"not null;foreignkey"`
}  

// POST /status
// Create new status
func (status *Status) Create(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Add log
		return
	}
	db.Create(&status)
	
	c.JSON(http.StatusOK, gin.H{"data": status})

}

// GET /status/:id
// Find a status
func (status *Status) FindByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(status).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// UPDATE /status/:id
// Update a status
func (status *Status) Update(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).Update(status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			// Add log
		return	
	db.Model(&status).Updates(status)

	c.JSON(http.StatusOK, gin.H{"data": status})
}

// DELETE /status/:id
// Delete a status
func (status *Status) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(status).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
	db.Delete(&status)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }





