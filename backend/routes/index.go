package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/handlers"
	"github.com/heronhoga/eid-fitr-card/middlewares"
)

func CardRoutes(r *fiber.App, h *handlers.CardHandler) {
	api := r.Group("/api")

	api.Get("/cards", middlewares.ApplicationKeyMiddleware, h.GetCard)
	api.Post("/cards", middlewares.ApplicationKeyMiddleware, h.CreateCard)
}