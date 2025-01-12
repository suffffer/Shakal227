package app

import "github.com/gofiber/fiber/v2"

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
	// api := app.Group("/api")

}
