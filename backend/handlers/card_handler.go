package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
	"github.com/heronhoga/eid-fitr-card/services"
)

type CardHandler struct {
	cardService *services.CardService
}

func NewCardHandler(cardService *services.CardService) *CardHandler {
	return &CardHandler{cardService: cardService}
}

func (h *CardHandler) GetCard(ctx fiber.Ctx) error {
	//*request
	var getCardRequest models.GetCardRequest
	err := ctx.Bind().Body(&getCardRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if getCardRequest.CardID == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "card id cannot be empty",
		})
	}

	//*service
	response, err := h.cardService.GetCard(ctx, getCardRequest)

	fmt.Println(response)

	//*response
	return ctx.Status(200).JSON(fiber.Map{
		"message": getCardRequest.CardID,
	})
}

func (h *CardHandler) CreateCard(ctx fiber.Ctx) error {
	//request
	var createCardRequest models.CreateCardRequest
	err := ctx.Bind().Body(&createCardRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if createCardRequest.To == "" || createCardRequest.From == "" || createCardRequest.Title == "" || createCardRequest.Description == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "all fields are required",
		})
	}
	//service
	card_id, err := h.cardService.CreateCard(ctx, createCardRequest)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	//response
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"card_id": card_id,
	})
}