package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func CropInspectorRoutes(cropinspector fiber.Router) {
	ctx := context.Background()

	cropinspector.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		inspectors, err := queries.GetCropInspectors(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspectors)

	})
	cropinspector.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, err := queries.CreateCropInspector(ctx, db.CreateCropInspectorParams{})
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		return c.JSON("cropinspector")
	})
	cropinspector.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/code/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		return c.JSON("cropinspector")
	})
	cropinspector.Post("/code/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		return c.JSON("cropinspector")
	})
}
