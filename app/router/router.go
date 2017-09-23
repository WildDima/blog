package router

import (
	"blogy/app/controllers"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	home := controllers.Home{}
	r.GET("/", home.Index)
}
