package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jayashreeks/shorten-url-fiber-redis/routes"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {

	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
