package main

import (
	"fmt"
	"log"
	"os"

	controller "github.com/MicahAsowata/elsa/cmd/web/controllers/task"
	"github.com/MicahAsowata/elsa/internal/db/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const port = ":3000"

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	dbName := os.Getenv("DB_NAME")
	hostName := os.Getenv("HOST_NAME")
	hostPswd := os.Getenv("HOST_PSWD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", hostName, hostPswd, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(models.Task{})
	if err != nil {
		log.Fatal(err)
	}
	controller.TaskRoutes(app)
	log.Println("Speak, for thy servant heareth")
	app.Listen(port)

}
