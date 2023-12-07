package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func FarmRoutes(farm fiber.Router) {
	ctx := context.Background()

	farm.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		farms, err := queries.GetFarms(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farms)
	})
	farm.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddFarm(ctx, db.AddFarmParams{})
		return c.JSON(res)
	})
	farm.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		farm, err := queries.GeTFarm(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farm)
	})
	farm.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")

		var name map[string]string
		err = c.BodyParser(name)

		farm, err := queries.UpdateFarm(ctx, db.UpdateFarmParams{
			Name:          name["name"],
			FarmValue:     farm_value["farm_value"],
			AddressStreet: address_street["address_street"],
			AddressCity:   address_city["address_city"],
			AddressState:  address_state["address_state"],
			AddressZip:    address_zip["address_zip"],
			FarmID:        int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farm)
	})

	farm.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		farm, err := queries.DeleteFarm(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farm)
	})
}
