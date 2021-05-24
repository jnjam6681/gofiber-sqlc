// package todo

// import (
// 	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
// )

// type Service interface {
// 	InsertTodo(todo *postgres.Todo) (*postgres.Todo, error)
// }

// type service struct {
// 	repository repository
// }

// func NewService(r repository) Service {
// 	return &service{
// 		repository: r,
// 	}
// }

// func (s *service) InsertTodo(todo *postgres.Todo) (*postgres.Todo, error) {
// 	return s.repository.InsertTodo(todo)
// }
