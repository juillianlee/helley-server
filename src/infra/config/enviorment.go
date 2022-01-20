package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	MONGO_URL            = ""
	DATABASE             = ""
	PORT                 = 0
	ACCESS_TOKEN_SECRET  = ""
	REFRESH_TOKEN_SECRET = ""
)

func LoadEnviorments() {
	var err error

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil || PORT == 0 {
		PORT = 8080
	}
	MONGO_URL = os.Getenv("MONGO_URL")
	DATABASE = os.Getenv("DATABASE")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")

	if MONGO_URL == "" {
		log.Fatal("Envioriment MONGO_URL not defined")
	}

	if DATABASE == "" {
		log.Fatal("Envioriment DATABASE not defined")
	}

	if ACCESS_TOKEN_SECRET == "" {
		log.Fatal("Envioriment ACCESS_TOKEN_SECRET not defined")
	}

	if ACCESS_TOKEN_SECRET == "" {
		log.Fatal("Envioriment ACCESS_TOKEN_SECRET not defined")
	}

	if REFRESH_TOKEN_SECRET == "" {
		log.Fatal("Envioriment REFRESH_TOKEN_SECRET not defined")
	}

}
