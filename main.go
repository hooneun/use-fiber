package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hooneun/use-fiber/router"
)

func main() {
	app := fiber.New()

	router.SetRoutes(app)

	app.Listen(":3333")

}
