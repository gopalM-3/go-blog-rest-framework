package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	fmt.Println("Env variables loaded!")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}