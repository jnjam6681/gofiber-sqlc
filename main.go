package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	postgres "github.com/jnjam6681/gofiber-sqlc/database/postgres/sqlc"
	"github.com/jnjam6681/gofiber-sqlc/pkg/routes"
	"github.com/jnjam6681/gofiber-sqlc/pkg/todo"

	_ "github.com/lib/pq"
)

const idleTimeout = 5 * time.Second

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s password=secret user=root sslmode=disable", "localhost", "todo_list"))
	if err != nil {
		panic(err)
	}

	result := postgres.NewRepo(db)
	repo := todo.NewRepository(result)

	todoService := todo.NewService(repo)

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api := app.Group("/api")
	routes.TodoRouter(api, todoService)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	db.Close()
	fmt.Println("Fiber was successful shutdown.")
}
