package main

import (
	"context"
	"log"
	"os"

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
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect MongoDB
	client, err := config.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// select database
	database := client.Database(os.Getenv("DB_NAME"))

	// layers
	cardRepo := repositories.NewCardRepository(database)
	cardService := services.NewCardService(cardRepo)
	cardHandler := handlers.NewCardHandler(cardService)

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// routes
	routes.CardRoutes(app, cardHandler)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}