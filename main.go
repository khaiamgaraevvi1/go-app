package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Salam, Dünya! 👋")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"author": "Xəyyam",
			"app":    "Go Fiber Project",
		})
	})

	app.Listen(":3000")
}
