package controllers

import (
	"blogy/app/models"
	"github.com/gin-gonic/gin"
)

type Posts struct {
	Application
}

func (c Posts) Create(context *gin.Context) {
	c.context = context

	Title := c.context.PostForm("title")
	Body := c.context.PostForm("body")

	post := models.Post{Title: Title, Body: Body}
	c.DB.Create(&post)
	c.status = 200
	c.body = JSON{"post": &post}
	c.render()
}
