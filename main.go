package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/app-info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":     "My Fiber App",
			"version": "1.0.0",
		})
	})

	app.Listen(":3000")
}
