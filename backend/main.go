package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/routes"
)

func main() {
    app := fiber.New()

    routes.UseRoutes(app)

    app.Listen(":3000")
}