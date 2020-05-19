package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	 _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/api"
	store "github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/db/mysql"
)

var r *gin.Engine

func main() {
	r := gin.Default()

	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
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
