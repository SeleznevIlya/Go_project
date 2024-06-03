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
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
