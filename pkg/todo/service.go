package todo

import (
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
)

type Service interface {
	InsertTodo(todo *postgres.Todo) (*postgres.Todo, error)
	ListTodos() (*[]postgres.Todo, error)
	GetTodo(id int64) (*postgres.Todo, error)
	DeleteTodo(todo *postgres.Todo) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertTodo(todo *postgres.Todo) (*postgres.Todo, error) {
	return s.repository.InsertTodo(todo)
}

func (s *service) ListTodos() (*[]postgres.Todo, error) {
	return s.repository.ListTodos()
}

func (s *service) GetTodo(id int64) (*postgres.Todo, error) {
	return s.repository.GetTodo(id)
}

func (s *service) DeleteTodo(todo *postgres.Todo) error {
	return s.repository.DeleteTodo(todo)
}
