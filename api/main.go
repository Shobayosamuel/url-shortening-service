package main

import (
	"github.com/Shobayosamuel/url-shortening-service/routes"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fibre.App) {
	app.Get("/:url", routes.ResolveURL)
	app.POST("/api/v1", routes.ShortenURL)
}


func main() {
	err := godotenv.Load()
	app := fibre.New()
}