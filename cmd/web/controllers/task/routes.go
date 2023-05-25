package controller

import (
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App) {
	baseRouteName := app.Group("/tarea")
	baseRouteName.Get("/new", TaskNew)
	baseRouteName.Get("/", TaskIndex)
	baseRouteName.Post("/", TaskCreate)
	baseRouteName.Get("/:id", TaskShow)
	baseRouteName.Get("/:id/edit", TaskEdit)
	baseRouteName.Post("/:id", TaskUpdate)
	baseRouteName.Get("/:id/delete", TaskDestroy)
}
