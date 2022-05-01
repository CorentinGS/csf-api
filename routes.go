package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func New() *fiber.App {
	// Create new app
	app := fiber.New()

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

	collection := mg.Db.Collection("weather")

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Printf("insert error: %s", err)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Successfully inserted weather",
		"success": true,
	})
}

func GetData(ctx *fiber.Ctx) error {
	var result Data

	col := mg.Db.Collection("weather")

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"createdat", -1}})
	findOptions.SetLimit(1)

	cursor, err := col.Find(context.TODO(), bson.D{}, findOptions)
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&result)
	}

	if err != nil {
		fmt.Println("FindOne() ERROR:", err)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully got the weather",
		"data":    result,
		"success": true,
	})
}
