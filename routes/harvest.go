package routes

import (
	"context"

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
			FarmID:   int32(farm),
			CropType: cropType,
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	harvest.Get("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		cropType := c.Params("type", "")
		harvests, err := queries.GetHarvests(ctx, cropType)
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
			FarmID:   farm,
			CropType: cropType,
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

		var body map[string]int32
		err = c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.DeleteHarvest(ctx, db.DeleteHarvestParams{CropType: cropType, FarmID: farm, HarvestYear: body["year"]})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)

	})
}

func farmAndType(c *fiber.Ctx) (farmId int32, cropType string, err error) {
	id, err := c.ParamsInt("farm")
	if err != nil {
		return 0, "", err
	}
	cropType = c.Params("type", "")
	if err != nil {
		return 0, "", err
	}
	return int32(id), cropType, nil
}
