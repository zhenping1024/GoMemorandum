package models

import (
	"github.com/jinzhu/gorm"
)
type User struct{
	gorm.Model
	UserName string `form:"username" gorm:"primary_key"`
	Password string`form:"password"`
	Todolist []Memorandum
}
type Memorandum struct{
	gorm.Model
	Title string `json:"title"`
	Content string `json:"context"`
	FinishFlag string `json:"flag"`
	Deadline string `json:"deadline gorm:type:datetime;default:'time.now()'"`
	UserId uint
}
