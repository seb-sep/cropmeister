package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CropInspectorRoutes(cropinspector fiber.Router) {

	cropinspector.Get("", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")

	})
	cropinspector.Post("", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Post("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
}
