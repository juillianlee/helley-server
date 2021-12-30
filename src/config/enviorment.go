package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGO_URL  = ""
	DATABASE   = ""
	PORT       = 0
	JWT_SECRET = ""
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

	if MONGO_URL == "" {
		log.Fatal("Envioriment MONGO_URL not defined")
	}

	if DATABASE == "" {
		log.Fatal("Envioriment DATABASE not defined")
	}

	if JWT_SECRET == "" {
		log.Fatal("Envioriment JWT_SECRET not defined")
	}

}
