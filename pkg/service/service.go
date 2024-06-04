package service

import (
	"github.com/SeleznevIlya/Go_project"
	"github.com/SeleznevIlya/Go_project/pkg/repository"
)

type Autorization interface {
	CreateUser(user Go_project.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list Go_project.TodoList) (int, error)
	GetAll(userId int) ([]Go_project.TodoList, error)
	GetById(userId, listId int) (Go_project.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input Go_project.UpdateTodoList) error
}

type TodoItem interface {
	Create(userId, listId int, item Go_project.TodoItem) (int, error)
	GetAll(userId, listId int) ([]Go_project.TodoItem, error)
	GetById(userId, itemId int) (Go_project.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input Go_project.UpdateTodoItem) error
}

type Service struct {
	Autorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Authorization),
		TodoList:     NewTodoListService(repos.TodoList),
		TodoItem:     NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
