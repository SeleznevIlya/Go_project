package repository

import (
	"github.com/SeleznevIlya/Go_project"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user Go_project.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
