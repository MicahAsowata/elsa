package main

import (
	"fmt"

	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pocketbase/dbx"
	"go.uber.org/zap"
)

func (base *base) Index(c *fiber.Ctx) error {
	notes := &[]models.Notes{}
	err := base.db.Select("id", "name", "body").All(notes)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	return c.Render("notes/index", fiber.Map{
		"Notes": notes,
		"Title": "🔥All Notes🔥",
	})
}

func (base *base) New(c *fiber.Ctx) error {
	return c.Render("notes/new", fiber.Map{
		"Title": "⚡New Note⚡",
	})
}

func (base *base) Create(c *fiber.Ctx) error {
	note := models.Notes{}
	err := c.BodyParser(&note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}

	err = base.db.Model(&note).Insert()
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}

	return c.Redirect("/")
}

func (base *base) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	note := &models.Notes{}
	err = base.db.Select("id", "name", "body").Model(id, note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	return c.Render("notes/show", fiber.Map{
		"Name":  note.Title,
		"Body":  note.Body,
		"ID":    note.ID,
		"Title": note.Title,
	})
}

func (base *base) Edit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	note := &models.Notes{}
	err = base.db.Select("id", "name", "body").Model(id, note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	return c.Render("notes/edit", fiber.Map{

		"Name":  note.Title,
		"Body":  note.Body,
		"ID":    note.ID,
		"Title": note.Title,
	})
}

func (base *base) Update(c *fiber.Ctx) error {
	note := models.Notes{}
	err := c.BodyParser(&note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	_, err = base.db.Update("notes", dbx.Params{
		"title": note.Title,
		"body":  note.Body,
	}, dbx.Between("id", c.Params("id"), c.Params("id"))).Execute()

	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}

	return c.Redirect(fmt.Sprintf("/%s", c.Params("id")))
}

func (base *base) Destroy(c *fiber.Ctx) error {
	_, err := base.db.Delete("notes", dbx.Between("id", c.Params("id"), c.Params("id"))).Execute()
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	return c.Redirect("/")
}
