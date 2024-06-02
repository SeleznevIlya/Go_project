package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
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
	return &Repository{}
}
