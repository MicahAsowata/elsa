package main

import (
	"fmt"

	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (app *application) TaskIndex(c *fiber.Ctx) error {
	return c.SendString("Index")
}

func (app *application) TaskNew(c *fiber.Ctx) error {
	return c.SendString("New")
}

func (app *application) TaskCreate(c *fiber.Ctx) error {
	task := models.Tasks{
		Name:      "Happy",
		Details:   "Happy Happy",
		Completed: true,
	}

	err := app.db.Model(&task).Insert()
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
		return c.SendString(string(err.Error()))
	}

	return c.SendString("Successful")
}

func (app *application) TaskShow(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Showing task %s", id))
}

func (app *application) TaskEdit(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Editing task %s", id))
}

func (app *application) TaskUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Updating task %s", id))
}

func (app *application) TaskDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Deleting task %s", id))
}
