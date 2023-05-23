package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Hello homies 🖖🏾🤏🏾")
}
func main() {
	app := fiber.New()

	app.Get("/", home)

	log.Println("Listening baby")
	app.Listen(":3000")
}
