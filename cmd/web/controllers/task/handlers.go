package controller

import (
	"fmt"
	"time"

	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var db = models.Db()

func TaskIndex(c *fiber.Ctx) error {
	return c.SendString("Index")
}

func TaskNew(c *fiber.Ctx) error {
	task := models.Task{
		ID:        uuid.New(),
		Name:      "Test",
		Details:   "Testing if it works",
		Start:     time.Now(),
		End:       time.Now().Add(time.Hour),
		Completed: false,
		Tag:       "Sweet Home",
	}

	db.Create(&task)
	return c.SendString("New")
}

func TaskCreate(c *fiber.Ctx) error {
	return c.SendString("Create")
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
