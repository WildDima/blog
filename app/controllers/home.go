package controllers

import (
	"blogy/app/models"
	"github.com/gin-gonic/gin"
)

type Home struct {
	Application
}

func (c Home) Index(context *gin.Context) {
	var posts []models.Post
	c.DB.Find(&posts)
	c.context = context
	c.status = 200
	c.body = JSON{"posts": &posts}
	c.render()
}
