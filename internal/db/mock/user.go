package mock

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

type UserMockRepository struct {
	DB map[string]*model.User
}

func (ur *UserMockRepository) GetByID(id string) (*model.User, error) {
	return ur.DB[id], nil
}
