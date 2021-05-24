package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	delivery "github.com/jnjam6681/gofiber-sqlc/delivery/http"
	"github.com/jnjam6681/gofiber-sqlc/usecase"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=secret user=root sslmode=disable", "todo_list"))
	if err != nil {
		panic(err)
	}

	result := postgres.NewRepo(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api := app.Group("/api")

	todoService := usecase.NewService(result)
	delivery.TodoRouter(api, todoService)

	app.Listen(":3000")
}
