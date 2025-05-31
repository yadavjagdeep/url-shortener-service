package main

import (
	"fmt"
	"log"

	"github.com/jagdeep/url-shortener-service/internal/api"
	"github.com/jagdeep/url-shortener-service/internal/repositories"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("URL shortner service starting")

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	repositories.InitMysqlClient("root:jagdeep@1234@tcp(localhost:3306)", "urlshortner")

	router := api.Router()
	api.SetupRoutes(router)
	router.Run(":8000")

}
