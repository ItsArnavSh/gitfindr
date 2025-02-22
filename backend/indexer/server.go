package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupServer() *fiber.App {
	app := fiber.New()

	app.Post("/search", func(c *fiber.Ctx) error {

		type RequestBody struct {
			Query string `json:"query"`
		}

		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		results := bm25(body.Query)

		return c.JSON(fiber.Map{
			"results": results,
		})
	})

	return app
}
