package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HarvestRoutes(harvest fiber.Router) {
	harvest.Get("", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Post("", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Get("/:farm/:type", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("type")
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		return c.JSON("harvest")
	})
	harvest.Put("/:farm/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
}
