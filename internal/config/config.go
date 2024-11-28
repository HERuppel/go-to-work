package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	Port                     = "0"
	SecretKey                []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port = os.Getenv("API_PORT")
	if err != nil {
		Port = "3333"
	}

	DatabaseConnectionString = fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_TYPE"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
