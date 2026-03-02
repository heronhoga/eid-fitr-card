package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

    // routes initial
    routes.CardRoutes(app, cardHandler)

    err = app.Listen(":8000")

	if err != nil {
		log.Fatal(err)
	}
}