package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hooneun/use-fiber/databases"
	"github.com/hooneun/use-fiber/models"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword make hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	db := databases.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}
