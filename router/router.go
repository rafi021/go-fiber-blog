package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/rafi021/go-fiber-blog/controller"
)

// Set Up Routes setup router API

func SetUpRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api/v1", logger.New())

	// Routes
	api.Get("/health", controller.HealthCheck)

	// Auth
	// auth := api.Group("/auth")
	// auth.Post("/register", Register)
	// auth.Post("/login", Login)

	// // Product routes
	// api.Get("/products", GetProducts)
	// api.Get("/products/:id", GetProduct)
	// api.Post("/products", CreateProduct)
	// api.Put("/products/:id", UpdateProduct)
	// api.Delete("/products/:id", DeleteProduct)

	// // User routes
	// api.Get("/users", GetUsers)
	// api.Get("/users/:id", GetUser)
	// api.Post("/users", CreateUser)
	// api.Put("/users/:id", UpdateUser)
	// api.Delete("/users/:id", DeleteUser)

}
