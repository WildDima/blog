package controllers

import "github.com/gin-gonic/gin"

type Home struct {
	Application
}

func (c Home) Index(context *gin.Context) {
	c.context = context
	c.status = 200
	c.body = JSON{"message": "pong"}
	c.render()
}
