package repository

import (
	"context"
	"errors"
	"fmt"
	"goServer/app"
	"goServer/models"
	"log"

	"github.com/jinzhu/gorm"
)

type ToDoRepo struct {
	DbConn *gorm.DB
}

func NewToDoRepo(Conn *gorm.DB) app.ToDoRepoInterface {
	return &ToDoRepo{
		DbConn: Conn,
	}
}

func (r *ToDoRepo) CreateToDo(ctx context.Context, title string, description string, priority, dueDate, tag string) (*models.ApiResponse, error) {
	t := &models.ToDo{
		Title:       title,
		Description: description,
		Tag:         tag,
		Priority:    priority,
		DueDate:     dueDate,
	}
	if err := r.DbConn.Create(t).Error; err != nil {
		return &models.ApiResponse{Msg: "failed"}, err
	}
	return &models.ApiResponse{Msg: "success"}, nil
}
func (r *ToDoRepo) UpdateToDoById(ctx context.Context, id, title, description, priority, dueDate, isCompleted, tag string) (*models.ApiResponse, error) {
	t := map[string]interface{}{
		"title":        title,
		"description":  description,
		"priority":     priority,
		"due_date":     dueDate,
		"tag":          tag,
		"is_completed": isCompleted,
	}
	fmt.Println(t)
	if rowsAffected := r.DbConn.Model(&models.ToDo{}).Where("id= ? ", id).Updates(t).RowsAffected; rowsAffected == 0 {
		return &models.ApiResponse{Msg: "failed"}, errors.New("not updated")
	}
	return &models.ApiResponse{Msg: "success"}, nil
}

func (r *ToDoRepo) GetToDos(ctx context.Context, limit int64, offset int64) (*models.ApiResponse, error) {
	t := &[]models.ToDo{}
	if err := r.DbConn.Model(&models.ToDo{}).Where("id>?", offset).Limit(limit).Scan(&t).Error; err != nil {
		log.Println(err)
		return &models.ApiResponse{Msg: "failed"}, err
	}
	return &models.ApiResponse{Msg: "success", ToDoList: t}, nil
}

func (r *ToDoRepo) DeleteToDoById(ctx context.Context, id string) (*models.ApiResponse, error) {
	t := &models.ToDo{}
	if rowsAffected := r.DbConn.Where("id= ? ", id).Delete(t).RowsAffected; rowsAffected == 0 {
		return &models.ApiResponse{Msg: "failed"}, errors.New("to not deleted")
	}
	return &models.ApiResponse{Msg: "success"}, nil
}
