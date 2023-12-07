package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func PurchaseRoutes(purchase fiber.Router) {
	ctx := context.Background()

	purchase.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		purchases, err := queries.GetPurchases(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchases)
	})
	purchase.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddPurchase(ctx, db.AddPurchaseParams{})
		return c.JSON(res)
	})
	purchase.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		purchase, err := queries.GetPurchase(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchase)
	})
	purchase.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		return c.JSON("purchase")
	})
	purchase.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		purchase, err := queries.DeletePurchase(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchase)
	})

}
