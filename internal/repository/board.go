package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

// BoardRepository implements board CRUD interface
type BoardRepository interface {
	GetByID(string) (*model.Board, error)
	GetByName(string) (*model.Board, error)
	Create(*model.Board, *gin.Context) (*model.Board, error)
	Update(string, *gin.Context) (*model.Board, error)
	Delete(string, *gin.Context) (*model.Board, error)
}