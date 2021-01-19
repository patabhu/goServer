package usecase

import (
	"context"
	"goServer/app"
	"goServer/models"
	"strconv"
)

type ToDoUsecase struct {
	formRepo app.ToDoRepoInterface
}

func NewToDoUsecase(repo app.ToDoRepoInterface) app.ToDoUsecaseInterface {
	return &ToDoUsecase{
		formRepo: repo,
	}
}
func (r *ToDoUsecase) CreateToDo(ctx context.Context, title, description, priority, dueDate, tag string) (*models.ApiResponse, error) {
	return r.formRepo.CreateToDo(ctx, title, description, priority, dueDate, tag)
}
func (r *ToDoUsecase) UpdateToDoById(ctx context.Context, id, title, description, priority, dueDate, isCompleted, tag string) (*models.ApiResponse, error) {
	// iID, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	return &models.ApiResponse{Msg: "failed"}, err
	// }
	return r.formRepo.UpdateToDoById(ctx, id, title, description, priority, dueDate, isCompleted, tag)
}
func (r *ToDoUsecase) GetToDos(ctx context.Context, limit string, offset string) (*models.ApiResponse, error) {
	limiti, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return &models.ApiResponse{Msg: "failed"}, err
	}
	offseti, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		return &models.ApiResponse{Msg: "failed"}, err
	}
	return r.formRepo.GetToDos(ctx, limiti, offseti)
}
func (r *ToDoUsecase) DeleteToDoById(ctx context.Context, id string) (*models.ApiResponse, error) {
	return r.formRepo.DeleteToDoById(ctx, id)
}
