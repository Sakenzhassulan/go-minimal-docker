package main

import (
	"github.com/Sakenzhassulan/go-minimal-docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
