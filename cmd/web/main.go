package main

import (
	"log"

	controller "github.com/MicahAsowata/elsa/cmd/web/controllers/task"
	"github.com/gofiber/fiber/v2"
)

const port = ":3000"

func main() {
	app := fiber.New()
	controller.TaskRoutes(app)
	log.Println("Speak, for thy servant heareth")
	app.Listen(port)

}
