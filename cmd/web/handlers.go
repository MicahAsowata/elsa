package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pocketbase/dbx"
	"go.uber.org/zap"
)

func (app *application) TaskIndex(c *fiber.Ctx) error {
	tasks := &[]models.Tasks{}
	err := app.db.Select("id", "name", "details", "completed").All(tasks)

	log.Println(err)
	log.Println(tasks)
	return c.Render("tasks/index", fiber.Map{
		"Message": tasks,
		"Title":   "All Tasks",
	})
}

func (app *application) TaskNew(c *fiber.Ctx) error {
	return c.Render("tasks/new", fiber.Map{
		"Message": "🐍",
		"Title":   "New Tasks",
	})
}

func (app *application) TaskCreate(c *fiber.Ctx) error {
	task := models.Tasks{
		Name:      "Joy ❤️❤️❤️",
		Details:   "We are Happy",
		Completed: false,
	}

	err := app.db.Model(&task).Insert()
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
		return c.SendString(string(err.Error()))
	}

	return c.SendString("Successful")
}

func (app *application) TaskShow(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	task := &models.Tasks{}
	err = app.db.Select("id", "name", "details", "completed").Model(id, task)
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	return c.Render("tasks/show", fiber.Map{
		"Name":      task.Name,
		"Details":   task.Details,
		"Completed": task.Completed,
		"ID":        task.ID,
		"Title":     task.Name,
	})
}

func (app *application) TaskEdit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	task := &models.Tasks{}
	err = app.db.Select("id", "name", "details", "completed").Model(id, task)
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	return c.Render("tasks/edit", fiber.Map{
		"Completed": task.Completed,
		"Name":      task.Name,
		"Details":   task.Details,
		"ID":        task.ID,
		"Title":     task.Name,
	})
}

func (app *application) TaskUpdate(c *fiber.Ctx) error {
	task := models.Tasks{}
	err := c.BodyParser(&task)
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	_, err = app.db.Update("tasks", dbx.Params{
		"name":    task.Name,
		"details": task.Details,
	}, dbx.Between("id", c.Params("id"), c.Params("id"))).Execute()

	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	c.Redirect(fmt.Sprintf("/tarea/%s", c.Params("id")), http.StatusOK)
	return nil
}

func (app *application) TaskDestroy(c *fiber.Ctx) error {
	_, err := app.db.Delete("tasks", dbx.Between("id", c.Params("id"), c.Params("id"))).Execute()
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}
	c.Redirect("/tarea/", http.StatusOK)
	return nil
}
