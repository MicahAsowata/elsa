package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const baseRouteName = "tarea"

func TaskRoutes(app *fiber.App) {
	app.Get(fmt.Sprintf("/%s/", baseRouteName), TaskIndex)
	app.Get(fmt.Sprintf("/%s/new", baseRouteName), TaskNew)
	app.Post(fmt.Sprintf("/%s/", baseRouteName), TaskCreate)
	app.Get(fmt.Sprintf("/%s/:id", baseRouteName), TaskShow)
	app.Get(fmt.Sprintf("/%s/:id/edit", baseRouteName), TaskEdit)
	app.Post(fmt.Sprintf("/%s/:id", baseRouteName), TaskUpdate)
	app.Get(fmt.Sprintf("/%s/:id/delete", baseRouteName), TaskDestroy)
}
