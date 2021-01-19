package app

import (
	"context"
	"goServer/models"
)

type ToDoUsecaseInterface interface {
	CreateToDo(ctx context.Context, title, description, priority, dueDate, tag string) (*models.ApiResponse, error)
	UpdateToDoById(ctx context.Context, id, title, description, priority, dueDate, isCompleted, tag string) (*models.ApiResponse, error)
	GetToDos(ctx context.Context, limit string, offset string) (*models.ApiResponse, error)
	DeleteToDoById(ctx context.Context, id string) (*models.ApiResponse, error)
}
