package service

import (
	"github.com/SeleznevIlya/Go_project"
	"github.com/SeleznevIlya/Go_project/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list Go_project.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]Go_project.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (Go_project.TodoList, error) {
	return s.repo.GetById(userId, listId)

}
