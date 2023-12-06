package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FarmerRoutes(farmer fiber.Router) {
	farmer.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})
	farmer.Post("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})

}
