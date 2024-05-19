package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	user := os.Getenv("user")
	pass := os.Getenv("pass")
	dbName := os.Getenv("dbname")
	port := os.Getenv("port")

	url := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
