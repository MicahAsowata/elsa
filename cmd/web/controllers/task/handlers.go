package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func TaskIndex(c *fiber.Ctx) error {
	return c.SendString("Index")
}

func TaskNew(c *fiber.Ctx) error {
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
