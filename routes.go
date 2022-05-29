package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/patrickmn/go-cache"
	"time"
)

func New() *fiber.App {
	// Create new app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, *",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	// Api group
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	// Auth
	api.Post("/data", PostData)
	api.Get("/data", GetData)

	return app
}

func PostData(ctx *fiber.Ctx) error {
	var data Data

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedAt = time.Now().UTC()

	c.Set("data", data, cache.NoExpiration)

	return ctx.JSON(fiber.Map{
		"message": "Successfully inserted weather",
		"success": true,
	})
}

func GetData(ctx *fiber.Ctx) error {
	result, found := c.Get("data")

	if found {
		return ctx.JSON(fiber.Map{
			"message": "Successfully retrieved weather",
			"success": true,
			"data":    result,
		})
	} else {
		return ctx.JSON(fiber.Map{
			"message": "No data found",
			"success": false,
			"data":    nil,
		})
	}
}
