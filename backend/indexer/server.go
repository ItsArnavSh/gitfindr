package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupServer() *fiber.App {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Change to frontend URL if needed
		AllowMethods: "POST, OPTIONS",
		AllowHeaders: "Content-Type",
	}))

	// Search route
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

		results := bm25(body.Query) // Your search logic

		return c.JSON(fiber.Map{
			"results": results,
		})
	})

	// Indexing route
	app.Post("/index", func(c *fiber.Ctx) error {
		type IndexRequest struct {
			RepoURL string `json:"repo_url"`
		}

		var body IndexRequest
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		index(body.RepoURL) // Run indexing in the background

		return c.JSON(fiber.Map{
			"message": "completed",
		})
	})

	return app
}
