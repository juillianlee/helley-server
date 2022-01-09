package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGO_URL            = ""
	DATABASE             = ""
	PORT                 = 0
	ACESSS_TOKEN_SECRET  = ""
	REFRESH_TOKEN_SECRET = ""
)

func LoadEnviorments() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil || PORT == 0 {
		PORT = 8080
	}
	MONGO_URL = os.Getenv("MONGO_URL")
	DATABASE = os.Getenv("DATABASE")
	ACESSS_TOKEN_SECRET = os.Getenv("ACESSS_TOKEN_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")

	if MONGO_URL == "" {
		log.Fatal("Envioriment MONGO_URL not defined")
	}

	if DATABASE == "" {
		log.Fatal("Envioriment DATABASE not defined")
	}

	if ACESSS_TOKEN_SECRET == "" {
		log.Fatal("Envioriment ACESSS_TOKEN_SECRET not defined")
	}

	if ACESSS_TOKEN_SECRET == "" {
		log.Fatal("Envioriment ACESSS_TOKEN_SECRET not defined")
	}

	if REFRESH_TOKEN_SECRET == "" {
		log.Fatal("Envioriment REFRESH_TOKEN_SECRET not defined")
	}

}
