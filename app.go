package main

import (
	books "fiber-book-app/controllers"
	"fiber-book-app/database"
	"fiber-book-app/helpers/env"
	"fiber-book-app/middlewares"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	fmt.Println("Environment: [" + env.Values.Env + "]")

	// connect to database
	db.ConnectToDatabase()

	// Create new Fiber instance
	app := fiber.New()

	// Set development middlewares
	if env.IsDevelopment() {
		// use logger
		app.Use(logger.New())
	}

	// Setup Panic Recover middleware
	app.Use(recover.New(recover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	}))

	// Allow CORS
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Fiber Book App",
		})
	})

	// Register routes
	app.Get("/books", books.AllBooks)
	app.Post("/books", books.AddBook)
	app.Get("/books/:id", middlewares.LoadBookParam, books.GetBook)
	app.Put("/books/:id", middlewares.LoadBookParam, books.UpdateBook)
	app.Delete("/books", books.DeleteAllBooks)

	log.Fatal(app.Listen("0.0.0.0:" + env.Values.AppPort))
}
