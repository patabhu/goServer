package models

import "github.com/jinzhu/gorm"

//api_repsonse
type ApiResponse struct {
	Msg      string  `json:"msg,omitempty"`
	ToDoList *[]ToDo `json:"to_dos,omitempty"`
}

//to_dos model for database
type ToDo struct {
	gorm.Model
	Title       string `json:"title,omitempty" gorm:"type:text;"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Tag         string `json:"tag,omitempty"`
	IsCompleted string `gorm:"type:boolean;default:false;" json:"is_completed,omitempty"`
}
