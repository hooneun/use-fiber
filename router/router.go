package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hooneun/use-fiber/handler"
)

// SetRoutes setting router
func SetRoutes(app *fiber.App) {
	// api := app.Group("/api", logger.New())
	api := app.Group("/api")
	api.Get("/", handler.Hello)
}
