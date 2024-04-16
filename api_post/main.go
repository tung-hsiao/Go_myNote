package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Struct to represent a simple item
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item // Slice to store items

func main() {
	// Create a new Fiber instance
	app := fiber.New()
	app.Use(compress.New()) // 啟用壓縮
	app.Use(cors.New())   // 啟用CORS
	app.Use(logger.New()) // 啟用日誌

	// Route to get all items
	app.Get("/api/items", func(c *fiber.Ctx) error {
		return c.JSON(items)
	})

	// Route to create a new item
	app.Post("/api/items", func(c *fiber.Ctx) error {
		// Parse JSON body into a new item
		var newItem Item
		c.BodyParser(&newItem)

		// Assign a unique ID (for simplicity, incrementing from len(items))
		newItem.ID = len(items) + 1

		// Add the new item to the slice
		items = append(items, newItem)

		// Return the created item as JSON response
		return c.Status(fiber.StatusCreated).JSON(newItem)
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
