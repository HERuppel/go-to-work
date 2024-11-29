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

	DatabaseConnectionString = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_TYPE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
