package todo

import (
	"context"

	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
)

type Repository interface {
	InsertTodo(todo *postgres.Todo) (*postgres.Todo, error)
	ListTodos() (*[]postgres.Todo, error)
	GetTodo(id int64) (*postgres.Todo, error)
	DeleteTodo(todo *postgres.Todo) error
	UpdateTodo(todo *postgres.Todo) (*postgres.Todo, error)
}

type repository struct {
	repo *postgres.Repo
}

func NewRepository(repo *postgres.Repo) Repository {
	return &repository{repo: repo}
}

func (r *repository) InsertTodo(todo *postgres.Todo) (*postgres.Todo, error) {
	result, err := r.repo.CreateTodo(context.Background(), todo.Name)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) UpdateTodo(todo *postgres.Todo) (*postgres.Todo, error) {
	arg := postgres.UpdateTodoParams{
		ID:       todo.ID,
		Complete: todo.Complete,
	}

	result, err := r.repo.UpdateTodo(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) ListTodos() (*[]postgres.Todo, error) {
	result, err := r.repo.ListTodos(context.Background())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetTodo(id int64) (*postgres.Todo, error) {
	result, err := r.repo.GetTodo(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) DeleteTodo(todo *postgres.Todo) error {
	err := r.repo.DeleteTodo(context.Background(), todo.ID)
	if err != nil {
		return err
	}
	return nil
}
