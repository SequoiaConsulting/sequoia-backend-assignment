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
	users := router.Group("/user")
	users.GET("/:id", uApi.getUserByID)
	// users.GET("/:id", uApi.getUserByName)
	users.POST("/", uApi.createUser)
	// users.PUT("/:id", uApi.updateUser)
	// users.DELETE("/:id", uApi.deleteUser)
}

func NewUserAPI(ur repository.UserRepository) *UserAPI {
	return &UserAPI{ur}
}

func (uApi *UserAPI) createUser(c *gin.Context) {
	userCore := core.NewUserCore(uApi.repo)
	user := &model.User{}
	c.Bind(user)

	user, err := userCore.Create(user)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, user)
}

func (uApi *UserAPI) getUserByID(c *gin.Context) {
	id := c.Param("id")

	userCore := core.NewUserCore(uApi.repo)

	resp, err := userCore.GetByID(id) // Check this
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

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *UserAPI) updateUser(c *gin.Context) {
// 	id := c.Param("id")

// 	userCore := core.NewUserCore(uApi.ur)

// 	resp, err := userCore.Update(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *UserAPI) deleteUser(c *gin.Context) {
// 	id := c.Param("id")

// 	userCore := core.NewUserCore(uApi.ur)

// 	resp, err := userCore.Delete(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *UserAPI) getUserByName(c *gin.Context) {
// 	name := c.Param("name")

// 	userCore := core.NewUserCore(uApi.ur)

// 	resp, err := userCore.GetByName(name)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }
