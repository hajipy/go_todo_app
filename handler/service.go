package handler

import (
	"context"

	"github.com/hajipy/go-todo-app-api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddTaskService ListTasksService RegisterUserService LoginService
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, password string) (string, error)
}
