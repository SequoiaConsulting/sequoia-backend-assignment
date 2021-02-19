package api

import (
	"fmt"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

type StatusAPI struct {
	repo repository.StatusRepository
}

// AddRoutes adds the status routes to the app.
func (sApi *StatusAPI) AddRoutes(router *gin.Engine) {
	statuses := router.Group("/status")
	statuses.GET("/get/id/:id", sApi.getStatusByID)
	statuses.GET("/name/:name", sApi.getStatusByName)
	statuses.POST("/", sApi.createStatus)
	statuses.PUT("/update/:id", sApi.updateStatus)
	statuses.DELETE("/delete/id/:id", sApi.deleteStatus)
}

func NewStatusAPI(ur repository.StatusRepository) *StatusAPI {
	return &StatusAPI{ur}
}

func (sApi *StatusAPI) createStatus(c *gin.Context) {
	statusCore := core.NewStatusCore(sApi.repo)
	status := &model.Status{}
	c.Bind(status)
	resp, err := statusCore.Create(status)
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

func (sApi *StatusAPI) getStatusByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("Hello")

	statusCore := core.NewStatusCore(sApi.repo)

	resp, err := statusCore.GetByID(id)
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

func (sApi *StatusAPI) getStatusByName(c *gin.Context) {
	id := c.Param("name")
	fmt.Printf("%s", id)

	statusCore := core.NewStatusCore(sApi.repo)

	resp, err := statusCore.GetByName(id)
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


func (sApi *StatusAPI) updateStatus(c *gin.Context) {
	statusCore := core.NewStatusCore(sApi.repo)
	status := &model.Status{}
	c.Bind(status)
	id := c.Param("id")
	resp, err := statusCore.Update(id, status)
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


func (sApi *StatusAPI) deleteStatus(c *gin.Context) {
	statusCore := core.NewStatusCore(sApi.repo)
	status := &model.Status{}
	c.Bind(status)
	
	resp, err := statusCore.Delete(status)
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
