package routes

import (
	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	"github.com/jnjam6681/gofiber-sqlc/pkg/todo"
)

func TodoRouter(app fiber.Router, service todo.Service) {
	app.Post("/todo", addTodo(service))
	app.Get("/todos", listTodos(service))
	app.Get("/todo/:id", getTodo(service))
	app.Delete("/todo", deleteTodo(service))
	app.Patch("/todo", updateTodo(service))
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

func listTodos(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListTodos()

		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  err,
		})
	}
}

func getTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		// id, err := strconv.Atoi(paramsId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse id",
			})
		}

		// var requestBody postgres.Todo
		// err := c.BodyParser(&requestBody)
		// if err != nil {
		// 	_ = c.JSON(&fiber.Map{
		// 		"success": false,
		// 		"error":   err,
		// 	})
		// }

		result, err := service.GetTodo(int64(id))
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  err,
		})
	}
}

func deleteTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody postgres.Todo
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		err = service.DeleteTodo(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "deleted successfully",
		})
	}
}

func updateTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// type request struct {
		// 	ID       string `json:"id"`
		// 	Complete bool   `json:"complete"`
		// }

		var requestBody postgres.Todo
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, err := service.UpdateTodo(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  err,
		})
	}
}
