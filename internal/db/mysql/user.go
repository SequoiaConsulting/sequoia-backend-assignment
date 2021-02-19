package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// UserMySQLRepository stores database connection
type UserMySQLRepository struct {
	db *gorm.DB
}

// NewUserMySQLRepository accepts UserMySQLRepository and returns UserMySQLRepository
func NewUserMySQLRepository(db *gorm.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db}
}

// GetByID does a database query to get a user by ID
func (repo *UserMySQLRepository) GetByID(id string) (*model.User, error) {
	user := &model.User{}
	fmt.Printf("%s", id)
	if err := repo.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByName does a database query to get a user by name
func (repo *UserMySQLRepository) GetByName(id string) (*model.User, error) {
	user := &model.User{}
	if err := repo.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create does a database query to create a user
func (repo *UserMySQLRepository) Create(user *model.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update does a database query to update a user by ID
func (repo *UserMySQLRepository) Update(id string, user *model.User) error {
	if err := repo.db.Where("id = ?", id).Update(user).Error; err != nil {
		fmt.Printf("%v", user)
		return err
	}
	return nil
}

// Delete does a database query to delete a user by ID
func (repo *UserMySQLRepository) Delete(user *model.User) error {
	if err := repo.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
