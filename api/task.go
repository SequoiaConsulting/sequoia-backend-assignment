package api

import (
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

	"github.com/gin-gonic/gin"
)

type TaskAPI struct {
	repo repository.TaskRepository
}

// AddRoutes adds the task routes to the app.
func (uApi *TaskAPI) AddRoutes(router *gin.Engine) {
	tasks := router.Group("/task/")
	tasks.GET("/get/id/:id", uApi.getTaskByID)
	tasks.GET("/name/:name", uApi.getTaskByName)
	tasks.POST("/", uApi.createTask)
	tasks.PUT("/update/:id", uApi.updateTask)
	tasks.DELETE("/delete/id/:id", uApi.deleteTask)
}

func NewTaskAPI(ur repository.TaskRepository) *TaskAPI {
	return &TaskAPI{ur}
}

func (uApi *TaskAPI) createTask(c *gin.Context) {
	taskCore := core.NewTaskCore(uApi.repo)
	task := &model.Task{}
	c.Bind(task)
	resp, err := taskCore.Create(task)
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

func (uApi *TaskAPI) getTaskByID(c *gin.Context) {
	id := c.Param("id")

	taskCore := core.NewTaskCore(uApi.repo)

	resp, err := taskCore.GetByID(id)
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

func (uApi *TaskAPI) getTaskByName(c *gin.Context) {
	id := c.Param("name")

	taskCore := core.NewTaskCore(uApi.repo)

	resp, err := taskCore.GetByID(id)
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


func (uApi *TaskAPI) updateTask(c *gin.Context) {
	taskCore := core.NewTaskCore(uApi.repo)
	task := &model.Task{}
	c.Bind(task)
	resp, err := taskCore.Update(task)
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


func (uApi *TaskAPI) deleteTask(c *gin.Context) {
	taskCore := core.NewTaskCore(uApi.repo)
	task := &model.Task{}
	c.Bind(task)
	resp, err := taskCore.Delete(task)
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
