package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	"github.com/jnjam6681/gofiber-sqlc/pkg/routes"
	"github.com/jnjam6681/gofiber-sqlc/pkg/todo"

	_ "github.com/lib/pq"
)

type Handlers struct {
	Repo *postgres.Repo
}

func NewHandlers(repo *postgres.Repo) *Handlers {
	return &Handlers{Repo: repo}
}

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=secret user=root sslmode=disable", "todo_list"))
	if err != nil {
		panic(err)
	}

	result := postgres.NewRepo(db)
	r := todo.NewRepository(result)
	// handlers := NewHandlers(result)
	// print(handlers)

	todoService := todo.NewService(r)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api := app.Group("/api")
	routes.TodoRouter(api, todoService)

	app.Listen(":3000")
}
