package router

import (
	"blogy/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Route(r *gin.Engine, db *gorm.DB) {
	home := controllers.Home{}
	home.DB = db

	posts := controllers.Posts{}
	posts.DB = db

	r.GET("/", home.Index)
	r.POST("/posts", posts.Create)
}
