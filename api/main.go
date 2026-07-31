package main

import (
	"log"
	"os"

	"github.com/Deeksha-hub-05/Go_URLShortener/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go URL Shortener API is running 🚀")
	})

	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	if os.Getenv("APP_PORT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println(".env file not found")
		}
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
