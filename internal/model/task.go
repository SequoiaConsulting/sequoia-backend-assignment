package model

import (
	"github.com/jinzhu/gorm"
)

// Task model to store id, title, description, due date, assigned by, assigned to, statusID, boardID
type Task struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` //why?
	Title string `gorm:"size:256;not null"`
	Description string `gorm:"size:1024;not null"`
	DueDate *time.Time `gorm:"default:null"`
	AssignedBy string `gorm:"foreignkey"`
	AssignedTo string `gorm:"foreignkey"`
	StatusID int `gorm:"foreignkey"`
	BoardID int `gorm:"foreignkey"`
}

// POST /task
// Create new task
func (task *Task) Create(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Add log
		return
	}
	db.Create(&task)
	
	c.JSON(http.StatusOK, gin.H{"data": task})

}

// GET /task/:id
// Find a task
func (task *Task) FindByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(task).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UPDATE /task/:id
// Update a task
func (task *Task) Update(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).Update(task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			// Add log
		return	
	db.Model(&task).Updates(task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /task/:id
// Delete a task
func (task *Task) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(task).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
	db.Delete(&task)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }



