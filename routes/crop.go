package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

// type AddCropRequest struct {
// 	CropType           string  `json:"crop_type"`
// 	BasePrice          float64 `json:"base_price"`
// 	PhRangeWeight      float64 `json:"ph_range_weight"`
// 	PhRangeDesired     float64 `json:"ph_range_desired"`
// 	WaterNeededWeight  float64 `json:"water_needed_weight"`
// 	WaterNeededDesired float64 `json:"water_needed_desired"`
// 	SunRangeWeight     float64 `json:"sun_range_weight"`
// 	SunRangeDesired    float64 `json:"sun_range_desired"`
// 	Banned             bool    `json:"banned"`
// }

type AddCropRequest struct {
	CropType string         `json:"crop_type"`
	Values   MutCropRequest `json:"values"`
}

type MutCropRequest struct {
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

	// Retrieves all information of a crop of a particular {CropType}.
	// (Also returns calculated price.)
	crop.Get("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crops, err := queries.GetCropData(ctx, c.Params("type"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crops)
	})

	// Create a new crop.
	crop.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		var body AddCropRequest
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		crop, err := queries.AddCrop(ctx, db.AddCropParams{
			CropType:           body.CropType,
			BasePrice:          sql.NullFloat64{Float64: body.Values.BasePrice, Valid: true},
			PhRangeWeight:      sql.NullFloat64{Float64: body.Values.PhRangeWeight, Valid: true},
			PhRangeDesired:     sql.NullFloat64{Float64: body.Values.PhRangeDesired, Valid: true},
			WaterNeededWeight:  sql.NullFloat64{Float64: body.Values.WaterNeededWeight, Valid: true},
			WaterNeededDesired: sql.NullFloat64{Float64: body.Values.WaterNeededDesired, Valid: true},
			SunRangeWeight:     sql.NullFloat64{Float64: body.Values.SunRangeWeight, Valid: true},
			SunRangeDesired:    sql.NullFloat64{Float64: body.Values.SunRangeDesired, Valid: true},
			Banned:             sql.NullBool{Bool: body.Values.Banned, Valid: true},
		})

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})

	// Updates the mutable attributes of the crop of {CropType}.
	crop.Put("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		var body MutCropRequest
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		crop, err := queries.UpdateCrop(ctx, db.UpdateCropParams{
			CropType:           c.Params("type"),
			BasePrice:          sql.NullFloat64{body.BasePrice, true},
			PhRangeWeight:      sql.NullFloat64{body.PhRangeWeight, true},
			PhRangeDesired:     sql.NullFloat64{body.PhRangeDesired, true},
			WaterNeededWeight:  sql.NullFloat64{body.WaterNeededWeight, true},
			WaterNeededDesired: sql.NullFloat64{body.WaterNeededDesired, true},
			SunRangeWeight:     sql.NullFloat64{body.SunRangeWeight, true},
			SunRangeDesired:    sql.NullFloat64{body.SunRangeDesired, true},
			Banned:             sql.NullBool{body.Banned, true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})

	// Mark {CropType} as banned.
	crop.Delete("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		crop, err := queries.DeleteCrop(ctx, c.Params("type"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(crop)
	})
}
