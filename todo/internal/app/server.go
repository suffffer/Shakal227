package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suffffer/Shakal227/internal/handlers"
)

var err error

func StartServer() error {
	app := fiber.New()
	registerRoutes(app)
	err = app.Listen(":8000")
	if err != nil {
		return err
	}
	return nil
}

func registerRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/posts", handlers.CreatePost)
	api.Get("/posts", handlers.GetPosts)
	api.Get("/posts/:id", handlers.GetPosts)
	api.Put("/posts/:id", handlers.UpdatePost)
	api.Delete("/posts/:id", handlers.DeletePost)
}
