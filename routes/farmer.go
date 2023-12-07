package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func FarmerRoutes(farmer fiber.Router) {
	ctx := context.Background()

	farmer.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		farmers, err := queries.GetFarmer(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farmers)
	})
	farmer.Post("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddFarmer(ctx, db.AddFarmerParams{})
		return c.JSON(res)
	})

}
