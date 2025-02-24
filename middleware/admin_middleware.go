package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role").(string)

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied: Admins only"})
	}
	return c.Next()
}
