package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load(".env.api")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("MONGO_URI")
}
