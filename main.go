package main

import (
	"blogy/app/models"
	"blogy/app/router"
	"blogy/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Post{})
	r := gin.Default()
	router.Route(r)
	r.Run()
}
