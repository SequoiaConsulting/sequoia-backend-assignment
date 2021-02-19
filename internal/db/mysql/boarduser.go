package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardUserMySQLRepository stores database connection
type BoardUserMySQLRepository struct {
	db *gorm.DB
}

// NewBoardUserMySQLRepository accepts BoardUserMySQLRepository and returns BoardUserMySQLRepository
func NewBoardUserMySQLRepository(db *gorm.DB) *BoardUserMySQLRepository {
	return &BoardUserMySQLRepository{db}
}

// GetByID does a database query to get a boardBoardUser by ID
func (repo *BoardUserMySQLRepository) GetByID(id string) (*model.BoardUser, error) {
	boardUser := &model.BoardUser{}
	if err := repo.db.Where("id = ?", id).First(boardUser).Error; err != nil {
		return nil, err
	}
	return boardUser, nil
}

// Create does a database query to create a boardBoardUser
func (repo *BoardUserMySQLRepository) Create(boardUser *model.BoardUser) error {

	if err := repo.Create(boardUser); err != nil {
		return err
	}
	return nil
}

// Update does a database query to update a boardBoardUser by ID
func (repo *BoardUserMySQLRepository) Update(boardUser *model.BoardUser) error {
	if err := repo.db.Where("id = ?", boardUser.ID).Update(boardUser).Error; err != nil {
		return err
	}
	return nil
}

// Delete does a database query to delete a boardBoardUser by ID
func (repo *BoardUserMySQLRepository) Delete(boardUser *model.BoardUser) error {
	if err := repo.db.Delete(boardUser).Error; err != nil {
		return err
	}
	return nil
}
