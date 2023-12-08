package routes

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddCropRequest struct {
	CropType           string  `json:"cropType"`
	BasePrice          float64 `json:"basePrice"`
	PhRangeWeight      float64 `json:"phRangeWeight"`
	PhRangeDesired     float64 `json:"phRangeDesired"`
	WaterNeededWeight  float64 `json:"waterNeededWeight"`
	WaterNeededDesired float64 `json:"waterNeededDesired"`
	SunRangeWeight     float64 `json:"sunRangeWeight"`
	SunRangeDesired    float64 `json:"sunRangeDesired"`
	Banned             bool    `json:"banned"`
}

type MutCropRequest struct {
	BasePrice          float64 `json:"basePrice"`
	PhRangeWeight      float64 `json:"phRangeWeight"`
	PhRangeDesired     float64 `json:"phRangeDesired"`
	WaterNeededWeight  float64 `json:"waterReededWeight"`
	WaterNeededDesired float64 `json:"waterNeededDesired"`
	SunRangeWeight     float64 `json:"sunRangeWeight"`
	SunRangeDesired    float64 `json:"sunRangeDesired"`
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
		fmt.Printf("%s\n", c.Body())
		body := AddCropRequest{}
		err := c.BodyParser(&body)
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

	// Updates the mutable attributes of the crop of {CropType}.
	crop.Put("/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)

		body := MutCropRequest{}
		err := c.BodyParser(&body)
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
