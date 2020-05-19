package api

// import (
// 	"fmt"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
// )

// type BoardAPI struct {
// 	ur repository.BoardRepository
// }

// func NewBoardAPI(ur repository.BoardRepository) *BoardAPI {
// 	return &BoardAPI{ur}
// }

// func (bApi *BoardAPI) getBoardByID(c *gin.Context) {
// 	id := c.Param("id")

// 	boardCore := core.NewBoardCore(bApi.ur)

// 	resp, err := boardCore.GetByID(id)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (bApi *BoardAPI) AddRoutes(router *gin.Engine) {
// 	boards := router.Group("/board")
// 	boards.GET("/:id", bApi.getBoardByID)
// 	boards.POST("/", bApi.createBoard)
// 	boards.PUT("/:id", bApi.updateBoard)
// 	boards.DELETE("/:id", bApi.deleteBoard)
// 	boards.GET("/:name", bApi.getBoardByName)
// }

// func (bApi *BoardAPI) createBoard(c *gin.Context) {
// 	id := c.Param("id")

// 	boardCore := core.NewBoardCore(bApi.ur)

// 	resp, err := boardCore.Create(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (bApi *BoardAPI) updateBoard(c *gin.Context) {
// 	id := c.Param("id")

// 	boardCore := core.NewBoardCore(bApi.ur)

// 	resp, err := boardCore.Update(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (bApi *BoardAPI) deleteBoard(c *gin.Context) {
// 	id := c.Param("id")

// 	boardCore := core.NewBoardCore(bApi.ur)

// 	resp, err := boardCore.Delete(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (bApi *BoardAPI) getBoardByName(c *gin.Context) {
// 	name := c.Param("name")

// 	boardCore := core.NewBoardCore(bApi.ur)

// 	resp, err := boardCore.GetByName(name)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }
