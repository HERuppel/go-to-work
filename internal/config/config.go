package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	Port                     = "0"
	SecretKey                []byte
	SmtpHost                 = ""
	SmtpPort                 = 0
	SmtpEmail                = ""
	SmtpPassword             = ""
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

	SmtpHost = os.Getenv("SMTP_HOST")
	SmtpPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("ERROR_LOADING_SMTP_PORT")
		return
	}
	SmtpEmail = os.Getenv("SMTP_EMAIL")
	SmtpPassword = os.Getenv("SMTP_PASSWORD")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
