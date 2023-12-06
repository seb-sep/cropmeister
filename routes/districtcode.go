package routes

import (
	"github.com/gofiber/fiber/v2"
)

func DistrictCodeRoutes(districtcode fiber.Router) {
	districtcode.Get("", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Put("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Delete("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
}
