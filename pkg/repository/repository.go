package repository

import (
	"github.com/SeleznevIlya/Go_project"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user Go_project.User) (int, error)
	GetUser(username, password string) (Go_project.User, error)
}

type TodoList interface {
	Create(userId int, list Go_project.TodoList) (int, error)
	GetAll(userId int) ([]Go_project.TodoList, error)
	GetById(userId, listId int) (Go_project.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input Go_project.UpdateTodoList) error
}

type TodoItem interface {
	Create(listId int, list Go_project.TodoItem) (int, error)
	GetAll(userId, listId int) ([]Go_project.TodoItem, error)
	GetById(userId, itemId int) (Go_project.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input Go_project.UpdateTodoItem) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
