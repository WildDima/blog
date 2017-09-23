package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type JSON map[string]interface{}

type Application struct {
	status  int
	body    JSON
	context *gin.Context
	DB      *gorm.DB
}

func (c Application) render() {
	if c.status == 0 {
		c.status = 200
	}
	if c.body == nil {
		c.body = JSON{}
	}
	c.context.JSON(c.status, gin.H(c.body))
}
