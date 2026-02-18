package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/handlers"
)

func CardRoutes(r *fiber.App, h *handlers.CardHandler) {
	r.Get("/card", h.GetCard)
}