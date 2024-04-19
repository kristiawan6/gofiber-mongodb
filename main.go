package main

import (
	"gofiber-mongodb/src/routes"
	config "gofiber-mongodb/src/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()
	routes.Router(app)

	app.Listen(":8080")
}
