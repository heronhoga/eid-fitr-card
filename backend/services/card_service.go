package services

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
	"github.com/heronhoga/eid-fitr-card/repositories"
)

type CardService struct {
	cardRepository *repositories.CardRepository
}

func NewCardService(cardRepository *repositories.CardRepository) *CardService {
	return &CardService{cardRepository: cardRepository}
}

func (s *CardService) GetCard(ctx fiber.Ctx, request models.GetCardRequest) (models.GetCardResponse, error) {
	return models.GetCardResponse{}, nil
}

// func (s *CardService) CreateCard(ctx fiber.Ctx,) error {

// }