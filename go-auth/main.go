package main

import (
	"github.com/kaniusii/museum/database"
	"github.com/kaniusii/museum/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	port := ":8880"

	app.Listen(port)
}
