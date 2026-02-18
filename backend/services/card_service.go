package services

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
)

func GetCard(ctx fiber.Ctx, request models.CardRequest) (models.CardResponse, error) {
	return models.CardResponse{}, nil
}