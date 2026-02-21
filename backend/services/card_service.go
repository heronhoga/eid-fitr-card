package services

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
	"github.com/heronhoga/eid-fitr-card/repositories"
	"github.com/heronhoga/eid-fitr-card/utils"
)

type CardService struct {
	cardRepository *repositories.CardRepository
}

func NewCardService(cardRepository *repositories.CardRepository) *CardService {
	return &CardService{cardRepository: cardRepository}
}

func (s *CardService) GetCard(ctx fiber.Ctx, request models.GetCardRequest) (models.GetCardResponse, error) {

	card, err := s.cardRepository.GetCard(ctx, request)

	if err != nil {
		fmt.Println(err)
		return models.GetCardResponse{}, err
	}

	return card, nil
}

func (s *CardService) CreateCard(ctx fiber.Ctx, request models.CreateCardRequest) (string, error) {
	//generate card id

	newCard := &models.Card{
		CardID: utils.GenerateCardID(),
		To: request.To,
		From: request.From,
		Title: request.Title,
		Description: request.Description,
		Type: request.Type,
	}

	card_id, err := s.cardRepository.CreateCard(ctx, newCard)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return card_id, nil
}