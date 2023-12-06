package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FarmRoutes(farm fiber.Router) {
	farm.Get("", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Post("", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	farm.Delete("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
}
