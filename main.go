package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hooneun/use-fiber/databases"
	"github.com/hooneun/use-fiber/router"
)

func main() {
	app := fiber.New()

	databases.ConnectDB()

	router.SetRoutes(app)

	log.Fatal(app.Listen(":3333"))
	// defer databases.DB.Close()
}
