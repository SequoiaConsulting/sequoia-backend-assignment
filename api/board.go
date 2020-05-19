package api

import (
	"fmt"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

type BoardAPI struct {
	repo repository.BoardRepository
}

// AddRoutes adds the board routes to the app.
func (bApi *BoardAPI) AddRoutes(router *gin.Engine) {
	boards := router.Group("/board")
	boards.GET("/get/id/:id", bApi.getBoardByID)
	boards.GET("/name/:name", bApi.getBoardByName)
	boards.POST("/create", bApi.createBoard)
	boards.PUT("/update/id/:id", bApi.updateBoard)
	boards.DELETE("/delete/id/:id", bApi.deleteBoard)

	boardsId := boards.Group("/id/:board_id")
	boardsId.GET("/user/id/:user_id", bApi.boardUser)
	boardsId.GET("/status/name/:status_name", bApi.boardStatus)
}

func NewBoardAPI(ur repository.BoardRepository) *BoardAPI {
	return &BoardAPI{ur}
}

func (bApi *BoardAPI) createBoard(c *gin.Context) {
	boardCore := core.NewBoardCore(bApi.repo)
	board := &model.Board{}
	c.Bind(board)
	resp, err := boardCore.Create(board)
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

func (bApi *BoardAPI) getBoardByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("Hello")

	boardCore := core.NewBoardCore(bApi.repo)

	resp, err := boardCore.GetByID(id)
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

func (bApi *BoardAPI) getBoardByName(c *gin.Context) {
	id := c.Param("name")
	fmt.Printf("%s", id)

	boardCore := core.NewBoardCore(bApi.repo)

	resp, err := boardCore.GetByName(id)
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


func (bApi *BoardAPI) updateBoard(c *gin.Context) {
	boardCore := core.NewBoardCore(bApi.repo)
	board := &model.Board{}
	c.Bind(board)
	id := c.Param("id")
	resp, err := boardCore.Update(id, board)
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


func (bApi *BoardAPI) deleteBoard(c *gin.Context) {
	boardCore := core.NewBoardCore(bApi.repo)
	board := &model.Board{}
	c.Bind(board)
	
	resp, err := boardCore.Delete(board)
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

func (bApi *BoardAPI) boardStatus(c *gin.Context) {
	boardCore := core.NewBoardCore(bApi.repo)
	board := &model.Board{}
	user := &model.User{}
	c.Bind(board)
	c.Bind(user)
	id := c.Param("id")
	
	resp, err := boardCore.BoardStatus(id, board, user)
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

func (bApi *BoardAPI) boardUser(c *gin.Context) {
	boardCore := core.NewBoardCore(bApi.repo)
	board := &model.Board{}
	user := &model.User{}
	c.Bind(board)
	c.Bind(user)
	id := c.Param("id")
	
	resp, err := boardCore.BoardUser(id, board, user)
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