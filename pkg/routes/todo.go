package routes

import (
	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	"github.com/jnjam6681/gofiber-sqlc/pkg/todo"
)

type handler struct {
	service todo.Service
}

func TodoRouter(app fiber.Router, service todo.Service) {

	handler := &handler{
		service: service,
	}

	app.Post("/todo", handler.addTodo)
	app.Get("/todos", handler.listTodos)
	app.Get("/todo/:id", handler.getTodo)
	app.Delete("/todo", handler.deleteTodo)
	app.Patch("/todo", handler.updateTodo)
}

func (h *handler) addTodo(c *fiber.Ctx) error {
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

	result, err := h.service.InsertTodo(&requestBody)
	return c.JSON(&fiber.Map{
		"status": result,
		"error":  err,
	})
}

func (h *handler) listTodos(c *fiber.Ctx) error {
	result, err := h.service.ListTodos()

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

func (h *handler) getTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	// id, err := strconv.Atoi(paramsId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	result, err := h.service.GetTodo(int64(id))
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

func (h *handler) deleteTodo(c *fiber.Ctx) error {
	var requestBody postgres.Todo
	err := c.BodyParser(&requestBody)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	err = h.service.DeleteTodo(&requestBody)
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

func (h *handler) updateTodo(c *fiber.Ctx) error {
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

	result, err := h.service.UpdateTodo(&requestBody)
	return c.JSON(&fiber.Map{
		"status": result,
		"error":  err,
	})
}
