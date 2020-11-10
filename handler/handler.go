package handler

import "github.com/gofiber/fiber/v2"

// Hello "/" response
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Hello World!",
		"data":    nil,
	})
}
