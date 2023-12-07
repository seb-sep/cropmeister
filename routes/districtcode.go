package routes

import (
	"github.com/gofiber/fiber/v2"
)

func DistrictCodeRoutes(districtcode fiber.Router) {
	districtcode.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		districtcodes, err := queries.districtcodes(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcodes)
	})
	districtcode.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		districtcode, err := queries.districtcode(ctx, int32(id))
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
		var name map[string]string
		err = c.BodyParser(name)
		inspector, err := queries.UpdateDistrictCode(ctx, db.UpdateDistrictCodeParams{
			MaxWater: name["max_water"],
			MaxFert:  name["max_fert"],
			CropType: name["crop_type"],
			CodeID:   int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
	districtcode.Get("/crop/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		type, err := c.ParamsStr("type")
		districtcodes, err := queries.GetDistrictsWithCrop(ctx, type)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(districtcodes)
	})
	districtcode.Post("/crop/:type", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddDistrictCode(ctx, db.AddCropDistrictCodeParams{})
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
		res, _ := queries.AddDistrictCode(ctx, db.AddCropDistrictCodeParams{})
		return c.JSON(res)
	})
}
