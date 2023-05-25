package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/pocketbase/dbx"
	"go.uber.org/zap"
)

const port = ":3000"

type application struct {
	app    *fiber.App
	db     *dbx.DB
	logger *zap.Logger
}

func main() {
	logger, _ := zap.NewProduction()
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error", zap.Error(err))
	}
	hostName := os.Getenv("HOST_NAME")
	dbName := os.Getenv("DB_NAME")
	hostPswd := os.Getenv("HOST_PSWD")
	engine := html.New("./ui/html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, err := dbx.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", hostName, hostPswd, dbName))
	if err != nil {
		logger.Error("Error", zap.Error(err))
	}
	base := &application{
		app:    app,
		db:     db,
		logger: logger,
	}
	TaskRoutes(base)
	logger.Info("Speak, for thy servant heareth")
	base.app.Listen(port)

}
