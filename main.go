package main

import (
	"blogy/app/models"
	"blogy/app/router"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=blog_user dbname=blog_db sslmode=disable password=qwerty123")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Post{})
	r := gin.Default()
	router.Route(r)
	r.Run()
}
