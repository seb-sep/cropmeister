package routes

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddHarvestRequest struct {
	Quantity       int32   `json:"quantity"`
	HarvestYear    int32   `json:"harvestYear"`
	PhBase         float64 `json:"phBase"`
	PhFertilized   float64 `json:"phFertilized"`
	WaterRain      float64 `json:"waterRain"`
	WaterSprinkler float64 `json:"waterSprinkler"`
	Sun            int32   `json:"sun"`
	Price          float64 `json:"price"`
	Extinct        bool    `json:"extinct"`
}

type DeleteHarvestRequest struct {
	HarvestYear int32 `json:"harvestYear"`
}

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

		body := AddHarvestRequest{}
		err = c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.AddHarvest(ctx, db.AddHarvestParams{
			FarmID:         int32(farm),
			CropType:       cropType,
			Quantity:       sql.NullInt32{Int32: body.Quantity, Valid: true},
			HarvestYear:    body.HarvestYear,
			PhBase:         sql.NullFloat64{Float64: body.PhBase, Valid: true},
			PhFertilized:   sql.NullFloat64{Float64: body.PhFertilized, Valid: true},
			WaterRain:      sql.NullFloat64{Float64: body.WaterRain, Valid: true},
			WaterSprinkler: sql.NullFloat64{Float64: body.WaterSprinkler, Valid: true},
			Sun:            sql.NullInt32{Int32: body.Sun, Valid: true},
			Price:          sql.NullFloat64{Float64: body.Price, Valid: true},
			Extinct:        sql.NullBool{Bool: body.Extinct, Valid: true},
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

		body := AddHarvestRequest{}
		err = c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.UpdateHarvest(ctx, db.UpdateHarvestParams{
			FarmID:         farm,
			CropType:       cropType,
			HarvestYear:    body.HarvestYear,
			Quantity:       sql.NullInt32{Int32: body.Quantity, Valid: true},
			PhBase:         sql.NullFloat64{Float64: body.PhBase, Valid: true},
			PhFertilized:   sql.NullFloat64{Float64: body.PhFertilized, Valid: true},
			WaterRain:      sql.NullFloat64{Float64: body.WaterRain, Valid: true},
			WaterSprinkler: sql.NullFloat64{Float64: body.WaterSprinkler, Valid: true},
			Sun:            sql.NullInt32{Int32: body.Sun, Valid: true},
			Price:          sql.NullFloat64{Float64: body.Price, Valid: true},
			Extinct:        sql.NullBool{Bool: body.Extinct, Valid: true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(res)
	})

	harvest.Delete("/:farm/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		id, err := c.ParamsInt("farm")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fmt.Printf("%d\n", id)
		cropType := c.Params("type", "")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fmt.Printf("%s\n", cropType)
		fmt.Printf("%s\n", c.Body())
		body := DeleteHarvestRequest{}
		err = c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.DeleteHarvest(ctx, db.DeleteHarvestParams{CropType: cropType, FarmID: int32(id), HarvestYear: body.HarvestYear})
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
