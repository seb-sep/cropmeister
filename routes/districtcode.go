package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type DistrictCodeUpdateRequest struct {
	MaxWater float64 `json:"maxWater"`
	MaxFert  float64 `json:"maxFert"`
	CropType string  `json:"cropType"`
}

type DistrictCodeInspectorRequest struct {
	CodeID int32 `json:"codeId"`
}

func DistrictCodeRoutes(districtcode fiber.Router) {
	ctx := context.Background()

	districtcode.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		districtcodes, err := queries.GetDistrictCodes(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcodes)
	})

	districtcode.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		districtcode, err := queries.GetDistrictCode(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcode)
	})

	districtcode.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		districtcode, err := queries.DeleteDistrictCode(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcode)
	})

	districtcode.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddDistrictCode(ctx, db.AddDistrictCodeParams{})
		return c.JSON(res)
	})

	districtcode.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		data := DistrictCodeUpdateRequest{}
		err = c.BodyParser(&data)
		res, err := queries.UpdateDistrictCode(ctx, db.UpdateDistrictCodeParams{
			MaxWater: sql.NullFloat64{data.MaxWater, true},
			MaxFert:  sql.NullFloat64{data.MaxFert, true},
			CropType: sql.NullString{data.CropType, true},
			CodeID:   int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	districtcode.Get("/crop/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		cropType := c.Params("type", "")
		districtcodes, err := queries.GetDistrictsWithCrop(ctx, sql.NullString{String: cropType, Valid: true})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcodes)
	})

	districtcode.Post("/crop/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		data := DistrictCodeUpdateRequest{}
		err := c.BodyParser(&data)
		id, err := c.ParamsInt("id")
		res, err := queries.AddDistrictCode(ctx, db.AddDistrictCodeParams{
			MaxWater: sql.NullFloat64{data.MaxWater, true},
			MaxFert:  sql.NullFloat64{data.MaxFert, true},
			CropType: sql.NullString{data.CropType, true},
			CodeID:   int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	districtcode.Get("/inspector/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		districtcode, err := queries.GetDistrictsForInspector(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcode)
	})

	districtcode.Post("/inspector/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		usdaid, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		body := DistrictCodeInspectorRequest{}
		err = c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		res, _ := queries.AddDistrictToInspector(ctx, db.AddDistrictToInspectorParams{Usdaid: int32(usdaid), CodeID: body.CodeID})
		return c.JSON(res)
	})
}
