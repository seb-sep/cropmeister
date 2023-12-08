package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type FarmUpdateRequest struct {
	Name          string `json:"name"`
	FarmValue     int32  `json:"farm_value"`
	AddressStreet string `json:"address_street"`
	AddressCity   string `json:"address_city"`
	AddressState  string `json:"address_state"`
}

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
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		farm, err := queries.GetFarm(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farm)
	})

	farm.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var body FarmUpdateRequest
		err = c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		farm, err := queries.UpdateFarm(ctx, db.UpdateFarmParams{
			Name:          body.Name,
			FarmValue:     sql.NullInt32{Int32: body.FarmValue, Valid: true},
			AddressStreet: sql.NullString{String: body.AddressStreet, Valid: true},
			AddressCity:   sql.NullString{String: body.AddressCity, Valid: true},
			AddressState:  sql.NullString{String: body.AddressState, Valid: true},
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
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		farm, err := queries.DeleteFarm(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farm)
	})
}
