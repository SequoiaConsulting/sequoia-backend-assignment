package api

// import (
// 	"fmt"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
// )

// type BoardUserAPI struct {
// 	ur repository.BoardUserRepository
// }

// func NewBoardUserAPI(ur repository.BoardUserRepository) *BoardUserAPI {
// 	return &BoardUserAPI{ur}
// }

// func (buApi *BoardUserAPI) getBoardUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	boardUserCore := core.NewBoardUserCore(buApi.ur)

// 	resp, err := boardUserCore.GetByID(id)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (buApi *BoardUserAPI) AddRoutes(router *gin.Engine) {
// 	boardUsers := router.Group("/boardUser")
// 	boardUsers.GET("/:id", buApi.getBoardUserByID)
// 	boardUsers.POST("/", buApi.createBoardUser)
// 	boardUsers.PUT("/:id", buApi.updateBoardUser)
// 	boardUsers.DELETE("/:id", buApi.deleteBoardUser)
// 	boardUsers.GET("/:name", buApi.getBoardUserByName)
// }

// func (buApi *BoardUserAPI) createBoardUser(c *gin.Context) {
// 	id := c.Param("id")

// 	boardUserCore := core.NewBoardUserCore(buApi.ur)

// 	resp, err := boardUserCore.Create(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (buApi *BoardUserAPI) updateBoardUser(c *gin.Context) {
// 	id := c.Param("id")

// 	boardUserCore := core.NewBoardUserCore(buApi.ur)

// 	resp, err := boardUserCore.Update(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (buApi *BoardUserAPI) deleteBoardUser(c *gin.Context) {
// 	id := c.Param("id")

// 	// THis is not the best way.
// 	boardUserCore := core.NewBoardUserCore(buApi.ur)

// 	resp, err := boardUserCore.Detete(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (buApi *BoardUserAPI) getBoardUserByName(c *gin.Context) {
// 	name := c.Param("name")
// 	boardUserCore := core.NewBoardUserCore(buApi.ur)

// 	resp, err := boardUserCore.GetByName(name)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }
