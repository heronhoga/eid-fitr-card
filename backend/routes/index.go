package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/handlers"
)

func UseRoutes(r *fiber.App) {
	r.Get("/card", handlers.GetCard)
}