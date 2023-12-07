package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func HarvestRoutes(harvest fiber.Router) {
	harvest.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		harvests, err := queries.harvests(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(harvests)
	})
	harvest.Post("", func(c *fiber.Ctx) error {
		harvest, err := queries.harvest(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(harvest)
	})
	harvest.Get("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		
		type, err := c.ParamsStr("type")
		harvest, err := queries.
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		return c.JSON(harvest)
	})
	harvest.Put("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		
		return c.JSON("harvest")
	})
	harvest.Delete("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		type, err := c.ParamsStr("type")
		
		harvest, err := queries.
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(harvest)
	})
}
