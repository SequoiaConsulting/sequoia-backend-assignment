package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/api"
	store "github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/db/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var r *gin.Engine

func main() {
	r := gin.Default()

	db, err := gorm.Open("mysql", "dummy:password/trello?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}

	userStore := store.NewUserMySQLRepository(db)
	userApp := api.NewUserAPI(userStore)
	userApp.AddRoutes(r)

	err = r.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
