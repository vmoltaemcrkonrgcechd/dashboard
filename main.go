package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	log.Fatal(app.Listen(":80"))
}
