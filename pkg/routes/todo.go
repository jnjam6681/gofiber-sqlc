package routes

import (
	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	"github.com/jnjam6681/gofiber-sqlc/pkg/todo"
)

func TodoRouter(app fiber.Router, service todo.Service) {
	app.Post("/todos", addTodo(service))
}

func addTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// type request struct {
		// 	Name string `json:"name"`
		// }
		var requestBody postgres.Todo
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, err := service.InsertTodo(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  err,
		})
	}
}
