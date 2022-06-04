package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Problemas no arquivo \".env\", Err: %s", err)
	}
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DATABASE := os.Getenv("DATABASE")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	TIMEZONE := os.Getenv("TIMEZONE")
	dsn := "host=" + HOST + " user=" + USER + " password=" + PASSWORD + " dbname=" + DATABASE + " port=" + PORT + " sslmode=disable TimeZone=" + TIMEZONE
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
