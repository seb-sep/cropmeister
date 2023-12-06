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
		crops, err := queries.GetCropData(ctx, c.Params("type"))

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crops)
	})

	crop.Put("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		_, err := queries.UpdateCrop(ctx, db.UpdateCropParams{
			CropType: c.Params("type"),
		})

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON("crop")
	})

	// crop.Delete("/:type", func(c *fiber.Ctx) error {
	// 	queries := c.Locals("db").(*db.Queries)

	// 	_, err := queries.DeleteCrop(ctx, c.Params("type"))
	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}
	// 	return c.JSON("crop")
	// })
}
