package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/suffffer/Shakal227/internal/handlers"
)

var err error

func StartServer() error {
	app := fiber.New()

	// Разрешаем все CORS запросы
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Разрешаем запросы с любых адресов
		AllowMethods: "GET,POST,PUT,DELETE",                         // Разрешаем GET, POST, PUT, DELETE методы
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Разрешаем необходимые заголовки
	}))

	registerRoutes(app)

	err = app.Listen(":8000")
	if err != nil {
		return err
	}
	return nil
}

func registerRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.Register)

	api.Post("/posts", handlers.CreatePost)
	api.Get("/posts", handlers.GetPosts)
	api.Get("/posts/:id", handlers.GetPost)
	api.Put("/posts/:id", handlers.UpdatePost)
	api.Delete("/posts/:id", handlers.DeletePost)
}
