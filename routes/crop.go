package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CropRoutes(crop fiber.Router) {
	crop.Get("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	crop.Put("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	crop.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
}
