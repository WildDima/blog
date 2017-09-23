package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" gorm:"size:255"`
	Body  string `json:"body"`
}
