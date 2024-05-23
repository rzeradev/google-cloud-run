package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rzeradev/google-cloud-run/configs"
	"github.com/rzeradev/google-cloud-run/internal/handlers"
)

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	cfg, err := configs.LoadConfig(workdir)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	app := fiber.New()
	app.Get("/weather/:zipcode", handlers.GetWeather)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
