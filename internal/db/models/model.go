package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
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
	return db
}
