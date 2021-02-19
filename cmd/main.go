package main

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	 _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/api"
	store "github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/db/mysql"
	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
)

var r *gin.Engine

func main() {
	r := gin.Default()
	fmt.Println("Hello")
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
		if err != nil {
		log.Fatal(err.Error())
	}

	automigrateAll(db)
	
	userStore := store.NewUserMySQLRepository(db)
	userApp := api.NewUserAPI(userStore)
	userApp.AddRoutes(r)

	boardStore := store.NewBoardMySQLRepository(db)
	boardApp := api.NewBoardAPI(boardStore)
	boardApp.AddRoutes(r)

	err = r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func automigrateAll(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Board{})
}


