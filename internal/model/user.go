package model

import (
	"github.com/jinzhu/gorm"
)

// User model to store id, name, email, password, isAdmin
type User struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"size:256;not null"`
	Email string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	IsAdmin *bool `gorm:"default: false"`
}

// func generatePasswordHash( *User.Password) str string {
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

// }

// POST /user
// Create new user
func (user *User) Create(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Add log
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/:id
// Find a user
func (user *User) FindByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(user).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UPDATE /user/:id
// Update a user
func (user *User) Update(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).Update(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			// Add log
		return
	}
	db.Model(&user).Updates(user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /user/:id
// Delete a user
func (user *User) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
  
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(user).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  	// Add log
	  return
	}
	db.Delete(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }


