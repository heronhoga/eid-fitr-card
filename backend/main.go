package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/config"
	"github.com/heronhoga/eid-fitr-card/handlers"
	"github.com/heronhoga/eid-fitr-card/repositories"
	"github.com/heronhoga/eid-fitr-card/routes"
	"github.com/heronhoga/eid-fitr-card/services"
	"github.com/joho/godotenv"
)

func main() {
    // env initial
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // db initial
    db, err := config.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

    // layers initial
	cardRepo := repositories.NewCardRepository(db)
    cardService := services.NewCardService(cardRepo)
    cardHandler := handlers.NewCardHandler(cardService)
  
    app := fiber.New()

    // routes initial
    routes.CardRoutes(app, cardHandler)

    app.Listen(":3000")
}