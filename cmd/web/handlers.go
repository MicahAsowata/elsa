package main

import (
	"fmt"

	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func TaskIndex(c *fiber.Ctx) error {
	return c.SendString("Index")
}

func TaskNew(c *fiber.Ctx) error {
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

func TaskShow(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Showing task %s", id))
}

func TaskEdit(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Editing task %s", id))
}

func TaskUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Updating task %s", id))
}

func TaskDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Deleting task %s", id))
}
