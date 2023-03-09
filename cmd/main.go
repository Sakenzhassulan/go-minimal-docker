package main

import (
	"github.com/Sakenzhassulan/go-minimal-docker/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectDb()
	app := fiber.New()

	registerRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! greet mee")
	})

	app.Listen(":8080")
}
