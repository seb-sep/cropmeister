package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PurchaseRoutes(purchase fiber.Router) {

	purchase.Get("", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Post("", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Delete("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})

}
