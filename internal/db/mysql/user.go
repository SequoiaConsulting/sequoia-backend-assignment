package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// UserMySQLRepository stores database connection 
type UserMySQLRepository struct {
	conn *gorm.DB
}

// NewUserMySQLRepository accepts UserMySQLRepository and returns UserMySQLRepository
func NewUserMySQLRepository(conn *gorm.DB) *UserMySQLRepository {
	return &UserMySQLRepository{conn}
}

// GetByID does a database query to get a user by ID
func (ur *UserMySQLRepository) GetByID(id string) (*model.User, error) {
	db := ur.conn
	user := &model.User{}
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create does a database query to create a user
func (ur *UserMySQLRepository) Create(fields *model.User, c *gin.Context) (*model.User, error) {
	db := ur.conn
	user := &model.User{}
	// Validate input
	if err := c.ShouldBindJSON(user); err != nil {
		return nil, err
	}
	db.Create(fields)
	return user, nil
}

// Update does a database query to update a user by ID
func (ur *UserMySQLRepository) Update(id string, c *gin.Context) (*model.User, error) {
	db := ur.conn
	user := &model.User{}
	if err := db.Where("id = ?", c.Param("id")).Update(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete does a database query to delete a user by ID
func (ur *UserMySQLRepository) Delete(id string, c *gin.Context) (*model.User, error) {
	db := ur.conn
	user := &model.User{}
	if err := db.Where("id = ?", c.Param("id")).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
