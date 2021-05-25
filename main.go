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

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s password=secret user=root sslmode=disable", "localhost", "todo_list"))
	if err != nil {
		panic(err)
	}

	result := postgres.NewRepo(db)
	repo := todo.NewRepository(result)

	todoService := todo.NewService(repo)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api := app.Group("/api")
	routes.TodoRouter(api, todoService)

	app.Listen(":3000")
}
