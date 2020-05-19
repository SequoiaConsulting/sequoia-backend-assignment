package api

// import (
// 	"fmt"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/core"
// )

// type StatusAPI struct {
// 	ur repository.StatusRepository
// }

// func NewStatusAPI(ur repository.StatusRepository) *StatusAPI {
// 	return &StatusAPI{ur}
// }

// func (sApi *StatusAPI) getStatusByID(c *gin.Context) {
// 	id := c.Param("id")

// 	statusCore := core.NewStatusCore(sApi.ur)

// 	resp, err := statusCore.GetByID(id)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (sApi *StatusAPI) AddRoutes(router *gin.Engine) {
// 	statuses := router.Group("/status")
// 	statuses.GET("/:id", sApi.getStatusByID)
// 	statuses.GET("/:name", sApi.getStatusByName)
// 	statuses.POST("/", sApi.createStatus)
// 	statuses.PUT("/:id", sApi.updateStatus)
// 	statuses.DELETE("/:id", sApi.deleteStatus)
// }

// func (sApi *StatusAPI) createStatus(c *gin.Context) {
// 	id := c.Param("id")

// 	statusCore := core.NewStatusCore(sApi.ur)

// 	resp, err := statusCore.Create(id, c)

// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (sApi *StatusAPI) updateStatus(c *gin.Context) {
// 	id := c.Param("id")

// 	statusCore := core.NewStatusCore(sApi.ur)

// 	resp, err := statusCore.Update(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (sApi *StatusAPI) deleteStatus(c *gin.Context) {
// 	id := c.Param("id")

// 	statusCore := core.NewStatusCore(sApi.ur)

// 	resp, err := statusCore.Delete(id, c)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }

// func (sApi *StatusAPI) getStatusByName(c *gin.Context) {
// 	name := c.Param("name")

// 	statusCore := core.NewStatusCore(sApi.ur)

// 	resp, err := statusCore.GetByName(name)
// 	if err != nil {
// 		c.JSON(500, nil)
// 		return
// 	}

// 	fmt.Printf("%v", resp)

// 	c.JSON(200, resp)
// }
