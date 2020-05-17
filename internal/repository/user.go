package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// UserRepository implements user CRUD interface
type UserRepository interface {
	GetByID(string) (*model.User, error)
	Create(*model.User, *gin.Context) (*model.User, error)
	Update(string, *gin.Context) (*model.User, error)
	Delete(string, *gin.Context) (*model.User, error)
}