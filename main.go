package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/pocketbase/dbx"
	"go.uber.org/zap"
)

const port = "0.0.0.0:3000"

type base struct {
	base   *fiber.App
	db     *dbx.DB
	logger *zap.Logger
}

func main() {
	logger, _ := zap.NewProduction()
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error", zap.Error(err))
	}
	engine := html.New("./ui", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, _ := dbx.Open("mysql", os.Getenv("DSN"))
	if err := db.DB().Ping(); err != nil {
		logger.Error("Error", zap.Error(err))
	}
	base := &base{
		base:   app,
		db:     db,
		logger: logger,
	}
	err = routes(base)
	if err != nil {
		logger.Error("Error", zap.Error(err))
	}
	logger.Info("Speak, for thy servant heareth")
	base.base.Listen(port)
}
