package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) error {
	api := app.Group("/api")

	userRoutes(api)

	return nil
}
