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
	err := base.db.Select("id", "title", "body").From("notes").OrderBy("id DESC").All(notes)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "There are no notes to show you",
		})
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

	_, err = base.db.Insert("notes", dbx.Params{
		"title": note.Title,
		"body":  note.Body,
	}).Execute()
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "You've really important stuff to say, but the note is not creating",
		})
	}

	return c.Redirect("/")
}

func (base *base) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	note := &models.Notes{}
	err = base.db.Select("id", "title", "body").From("notes").Where(dbx.HashExp{"id": id}).One(&note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "We could not find that note.",
		})
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
	err = base.db.Select("id", "title", "body").Model(id, note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "It seems like the note doesn't want to be changed",
		})
	}
	return c.Render("notes/edit", fiber.Map{

		"Name":  note.Title,
		"Body":  note.Body,
		"ID":    note.ID,
		"Title": note.Title,
	})
}

func (base *base) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	note := models.Notes{}
	err := c.BodyParser(&note)
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
	}
	_, err = base.db.Update("notes", dbx.Params{
		"title": note.Title,
		"body":  note.Body,
	}, dbx.HashExp{"id": id}).Execute()

	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "The contents of the notes could not be changed.",
		})
	}

	return c.Redirect(fmt.Sprintf("/%s", c.Params("id")))
}

func (base *base) Destroy(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	_, err := base.db.Delete("notes", dbx.HashExp{"id": id}).Execute()
	if err != nil {
		base.logger.Error("Error", zap.Error(err))
		return c.Render("error", fiber.Map{
			"Title":   "⚠️ Error",
			"Message": "That note could not be deleted. We are loooking into it",
		})
	}
	return c.Redirect("/")
}
