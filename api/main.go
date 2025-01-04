package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shobayosamuel/url-shortening-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	
)

func setupRoutes(app *fibre.App) {
	app.Get("/:url", routes.ResolveURL)
	app.POST("/api/v1", routes.ShortenURL)
}


func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	app := fibre.New()

	app.Use(logger.New())
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}