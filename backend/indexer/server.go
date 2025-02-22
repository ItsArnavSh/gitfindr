package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupServer() *fiber.App {
	app := fiber.New()
	app.Get("/results", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Results!"})
	})
	return app
}
