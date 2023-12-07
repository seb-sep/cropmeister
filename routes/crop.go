package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddCropRequest struct {
	CropType           string  `json:"crop_type"`
	BasePrice          float64 `json:"base_price"`
	PhRangeWeight      float64 `json:"ph_range_weight"`
	PhRangeDesired     float64 `json:"ph_range_desired"`
	WaterNeededWeight  float64 `json:"water_needed_weight"`
	WaterNeededDesired float64 `json:"water_needed_desired"`
	SunRangeWeight     float64 `json:"sun_range_weight"`
	SunRangeDesired    float64 `json:"sun_range_desired"`
	Banned             bool    `json:"banned"`
}

func CropRoutes(crop fiber.Router) {
	ctx := context.Background()

	crop.Get("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crops, err := queries.GetCropData(ctx, c.Params("type"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crops)
	})

	crop.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		var body AddCropRequest
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		crop, err := queries.AddCrop(ctx, db.AddCropParams{
			CropType:           body.CropType,
			BasePrice:          sql.NullFloat64{Float64: body.BasePrice, Valid: true},
			PhRangeWeight:      sql.NullFloat64{Float64: body.PhRangeWeight, Valid: true},
			PhRangeDesired:     sql.NullFloat64{Float64: body.PhRangeDesired, Valid: true},
			WaterNeededWeight:  sql.NullFloat64{Float64: body.WaterNeededWeight, Valid: true},
			WaterNeededDesired: sql.NullFloat64{Float64: body.WaterNeededDesired, Valid: true},
			SunRangeWeight:     sql.NullFloat64{Float64: body.SunRangeWeight, Valid: true},
			SunRangeDesired:    sql.NullFloat64{Float64: body.SunRangeDesired, Valid: true},
			Banned:             sql.NullBool{Bool: body.Banned, Valid: true},
		})

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})

	crop.Put("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crop, err := queries.UpdateCrop(ctx, db.UpdateCropParams{
			CropType: c.Params("type"),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})

	crop.Delete("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crop, err := queries.DeleteCrop(ctx, c.Params("type"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})
}
