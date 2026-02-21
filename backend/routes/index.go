package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/handlers"
	"github.com/heronhoga/eid-fitr-card/middlewares"
)

func CardRoutes(r *fiber.App, h *handlers.CardHandler) {
	r.Get("/card", middlewares.ApplicationKeyMiddleware, h.GetCard)
	r.Post("/card", middlewares.ApplicationKeyMiddleware, h.CreateCard)
}