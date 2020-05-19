package api

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	repo repository.UserRepository
}

// AddRoutes adds the user routes to the app.
func (uApi *UserAPI) AddRoutes(router *gin.Engine) {
	users := router.Group("/user/")
	users.GET("/get/id/:id", uApi.getUserByID)
	users.GET("/name/:name", uApi.getUserByName)
	users.POST("/", uApi.createUser)
	users.PUT("/update/:id", uApi.updateUser)
	users.DELETE("/delete/id/:id", uApi.deleteUser)
}

func NewUserAPI(ur repository.UserRepository) *UserAPI {
	return &UserAPI{ur}
}

func (uApi *UserAPI) createUser(c *gin.Context) {
	userCore := core.NewUserCore(uApi.repo)
	user := &model.User{}
	c.Bind(user)
	resp, err := userCore.Create(user)
	if resp == nil {
		c.JSON(404, ErrorResponse{Message: "Not Found"})
		return
	}
	if err != nil {
		c.JSON(500, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

func (uApi *UserAPI) getUserByID(c *gin.Context) {
	id := c.Param("id")

	userCore := core.NewUserCore(uApi.repo)

	resp, err := userCore.GetByID(id)
	if resp == nil {
		c.JSON(404, ErrorResponse{Message: "Not Found"})
		return
	}
	if err != nil {
		c.JSON(500, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

func (uApi *UserAPI) getUserByName(c *gin.Context) {
	id := c.Param("id")

	userCore := core.NewUserCore(uApi.repo)

	resp, err := userCore.GetByID(id)
	if resp == nil {
		c.JSON(404, ErrorResponse{Message: "Not Found"})
		return
	}
	if err != nil {
		c.JSON(500, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}


func (uApi *UserAPI) updateUser(c *gin.Context) {
	userCore := core.NewUserCore(uApi.repo)
	user := &model.User{}
	c.Bind(user)
	resp, err := userCore.Update(user)
	if resp == nil {
		c.JSON(404, ErrorResponse{Message: "Not Found"})
		return
	}
	if err != nil {
		c.JSON(500, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, resp)
}


func (uApi *UserAPI) deleteUser(c *gin.Context) {
	userCore := core.NewUserCore(uApi.repo)
	user := &model.User{}
	c.Bind(user)
	resp, err := userCore.Delete(user)
	if resp == nil {
		c.JSON(404, ErrorResponse{Message: "Not Found"})
		return
	}
	if err != nil {
		c.JSON(500, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, resp)
}
