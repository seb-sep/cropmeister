package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

func CropInspectorRoutes(cropinspector fiber.Router) {
	ctx := context.Background()

	cropinspector.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		inspectors, err := queries.GetCropInspectors(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspectors)
	})

	cropinspector.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		res, _ := queries.AddCropInspector(ctx, db.AddCropInspectorParams{})
		return c.JSON(res)
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		inspector, err := queries.GetCropInspector(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
	cropinspector.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		var name map[string]string
		err = c.BodyParser(name)
		inspector, err := queries.UpdateCropInspector(ctx, db.UpdateCropInspectorParams{
			Name:   name["name"],
			Usdaid: int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
	cropinspector.Get("/code/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		inspector, err := queries.GetInspectorForDistrict(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
	cropinspector.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		inspector, err := queries.DeleteCropInspector(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
}
