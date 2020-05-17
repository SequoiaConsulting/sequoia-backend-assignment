package api

import (
	"fmt"

	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
)

type UserAPI struct {
	ur repository.UserRepository
}

func NewUserAPI(ur repository.UserRepository) *UserAPI {
	return &UserAPI{ur}
}

func (uapi *UserAPI) getUser(c *gin.Context) {
	id := c.Param("id")

	// THis is not the best way.
	userCore := core.NewUserCore(uapi.ur)

	resp, err := userCore.GetByID(id)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	fmt.Printf("%v", resp)

	c.JSON(200, resp)
}

func (uapi *UserAPI) AddRoutes(router *gin.Engine) {
	users := router.Group("/user")
	users.GET("/:id", uapi.getUser)
	users.POST("/", uapi.createUser)
	users.PUT("/:id", uapi.updateUser)
	users.DELETE("/:id", uapi.deleteUser)
}

func (uapi *UserAPI) createUser(c *gin.Context) {
	id := c.Param("id")

	// THis is not the best way.
	userCore := core.NewUserCore(uapi.ur)

	resp, err := userCore.GetByID(id)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	fmt.Printf("%v", resp)

	c.JSON(200, resp)
}

func (uapi *UserAPI) updateUser(c *gin.Context) {
	id := c.Param("id")

	// THis is not the best way.
	userCore := core.NewUserCore(uapi.ur)

	resp, err := userCore.GetByID(id)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	fmt.Printf("%v", resp)

	c.JSON(200, resp)
}

func (uapi *UserAPI) deleteUser(c *gin.Context) {
	id := c.Param("id")

	// THis is not the best way.
	userCore := core.NewUserCore(uapi.ur)

	resp, err := userCore.GetByID(id)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	fmt.Printf("%v", resp)

	c.JSON(200, resp)
}
