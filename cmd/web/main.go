package main

import (
	"log"

	"github.com/MicahAsowata/elsa/cmd/web/controller"
	"github.com/gofiber/fiber/v2"
)

const port = ":3000"

func main() {
	app := fiber.New()
	controller.TaskRoutes(app)
	log.Println("Listening baby")
	app.Listen(port)
}
