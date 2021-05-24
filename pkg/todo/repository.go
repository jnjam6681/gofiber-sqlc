package todo

import (
	"context"

	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
)

type Repository interface {
	InsertTodo(todo *postgres.Todo) (*postgres.Todo, error)
}

type repository struct {
	repo *postgres.Repo
}

func NewRepo(repo *postgres.Repo) Repository {
	return &repository{repo: repo}
}

func (r *repository) InsertTodo(todo *postgres.Todo) (*postgres.Todo, error) {
	ctx := context.Background()
	_, err := r.repo.CreateTodo(ctx, todo.Name)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
