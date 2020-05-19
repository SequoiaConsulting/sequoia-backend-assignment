package api

// import (
// 	"fmt"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
// )

// type TaskAPI struct {
// 	ur repository.TaskRepository
// }

// func NewTaskAPI(ur repository.TaskRepository) *TaskAPI {
// 	return &TaskAPI{ur}
// }

// func (uApi *TaskAPI) getTaskByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// THis is not the best way.
// 	taskCore := core.NewTaskCore(uApi.ur)

// 	resp, err := taskCore.GetByID(id)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *TaskAPI) AddRoutes(router *gin.Engine) {
// 	tasks := router.Group("/task")
// 	tasks.GET("/:id", uApi.getTaskByID)
// 	tasks.GET("/:name", uApi.getTaskByName)
// 	tasks.POST("/", uApi.createTask)
// 	tasks.PUT("/:id", uApi.updateTask)
// 	tasks.DELETE("/:id", uApi.deleteTask)
// }

// func (uApi *TaskAPI) createTask(c *gin.Context) {
// 	id := c.Param("id")

// 	taskCore := core.NewTaskCore(uApi.ur)

// 	resp, err := taskCore.Create(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *TaskAPI) updateTask(c *gin.Context) {
// 	id := c.Param("id")

// 	// THis is not the best way.
// 	taskCore := core.NewTaskCore(uApi.ur)

// 	resp, err := taskCore.Update(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *TaskAPI) deleteTask(c *gin.Context) {
// 	id := c.Param("id")

// 	// THis is not the best way.
// 	taskCore := core.NewTaskCore(uApi.ur)

// 	resp, err := taskCore.Delete(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (uApi *TaskAPI) getTaskByName(c *gin.Context) {
// 	name := c.Param("name")

// 	taskCore := core.NewTaskCore(uApi.ur)

// 	resp, err := taskCore.GetByName(name)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }
