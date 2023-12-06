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
	harvest.Get("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Put("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
}
