package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func CropRoutes(crop fiber.Router) {
	ctx := context.Background()

	crop.Get("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crops, err := queries.GetCropData(ctx, "test")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crops)
	})
	crop.Put("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	crop.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
}
