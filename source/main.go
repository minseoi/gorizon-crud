package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minseoi/gorizon/db"
	"github.com/minseoi/gorizon/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	e := echo.New()
	db.Initialize()
	routes.RegisterRoutes(e)

	port := getEnv("PORT", "8080")
	e.Logger.Fatal(e.Start(":" + port))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
