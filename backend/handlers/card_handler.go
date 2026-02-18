package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
	"github.com/heronhoga/eid-fitr-card/services"
)

func GetCard(ctx fiber.Ctx) error {
	//*requests
	var cardRequest models.CardRequest
	err := ctx.Bind().Body(&cardRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if cardRequest.CardID == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "card id cannot be empty",
		})
	}

	//*service
	response, err := services.GetCard(ctx, cardRequest)

	fmt.Println(response)
	
	//*response
	return ctx.Status(200).JSON(fiber.Map{
		"message": cardRequest.CardID,
	})
}
