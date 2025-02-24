package handlers

import (
	"yakz/config"
	"yakz/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(c *fiber.Ctx) error {

	db := config.GetDB()

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	hsahedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hsahedPassword)
	user.Role = "user"

	db.Create(&user)

	return c.JSON(fiber.Map{"Message": "User registered successfully "})
}
