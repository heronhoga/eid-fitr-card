package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func ApplicationKeyMiddleware(c fiber.Ctx) error {
	if c.Get("Application-Key") != os.Getenv("APPLICATION_KEY") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	
	return c.Next()
}