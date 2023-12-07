package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func HarvestRoutes(harvest fiber.Router) {
	ctx := context.Background()

	harvest.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		harvests, err := queries.GetAllHarvests(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(harvests)
	})

	harvest.Post("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		farm, cropType, err := farmAndType(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// TODO: Add more fields
		res, err := queries.AddHarvest(ctx, db.AddHarvestParams{
			FarmID:   sql.NullInt32{Int32: int32(farm), Valid: true},
			CropType: sql.NullString{String: cropType, Valid: true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	harvest.Get("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		cropType := c.Params("type", "")
		harvests, err := queries.GetHarvests(ctx, sql.NullString{String: cropType, Valid: true})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(harvests)
	})

	harvest.Put("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		farm, cropType, err := farmAndType(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		res, err := queries.UpdateHarvest(ctx, db.UpdateHarvestParams{
			FarmID:   sql.NullInt32{Int32: int32(farm), Valid: true},
			CropType: sql.NullString{String: cropType, Valid: true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(res)
	})

	harvest.Delete("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		farm, cropType, err := farmAndType(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.DeleteHarvest(ctx, db.DeleteHarvestParams{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)

	})
}

func farmAndType(c *fiber.Ctx) (farmId int, cropType string, err error) {
	farmId, err = c.ParamsInt("farm")
	if err != nil {
		return 0, "", err
	}
	cropType = c.Params("type", "")
	if err != nil {
		return 0, "", err
	}
	return farmId, cropType, nil
}
