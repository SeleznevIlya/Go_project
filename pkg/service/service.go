package service

import "github.com/SeleznevIlya/Go_project/pkg/repository"

type Autorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
