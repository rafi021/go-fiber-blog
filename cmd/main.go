package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/rafi021/go-fiber-blog/database"
)

func SetUpRoutes(app *fiber.App) error {
	api := app.Group("/api/v1")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Fiber APP is running......")
	})

	return nil
}

func main() {

	app := fiber.New()

	database.ConnectDB()

	// Enable CORS middleware
	app.Use(cors.New())

	// Set up routes
	SetUpRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8000"))

}
