package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pocketbase/dbx"
)

const port = ":3000"

type application struct {
	db *dbx.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load env")
	}
	hostName := os.Getenv("HOST_NAME")
	dbName := os.Getenv("DB_NAME")
	hostPswd := os.Getenv("HOST_PSWD")

	app := fiber.New()
	db, err := dbx.Open("mysql", fmt.Sprintf("mysql://%s:%s@tcp(localhost:3306)/%s", hostName, hostPswd, dbName))
	if err != nil {
		log.Println("Could not connect with the db")
	}

	base := application{
		db: db,
	}
	log.Print(base)
	log.Println("Speak, for thy servant heareth")
	app.Listen(port)

}
