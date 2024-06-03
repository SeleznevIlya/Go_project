package repository

import (
	"fmt"

	"github.com/SeleznevIlya/Go_project"
	"github.com/jmoiron/sqlx"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list Go_project.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}

func (r *TodoListRepository) GetAll(userId int) ([]Go_project.TodoList, error) {
	var lists []Go_project.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
							INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1`,
		todoListsTable, usersListsTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListRepository) GetById(userId, listId int) (Go_project.TodoList, error) {
	var list Go_project.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
							INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err

}
