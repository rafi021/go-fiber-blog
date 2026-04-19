package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/rafi021/go-fiber-blog/database"
	"github.com/rafi021/go-fiber-blog/router"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	// Enable CORS middleware
	app.Use(cors.New())

	// Set up routes
	router.SetUpRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8000"))

}
