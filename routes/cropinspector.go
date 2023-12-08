package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddCropInspectorRequest struct {
	Name   string `json:"name"`
	Usdaid int32  `json:"usdaid"`
}

type UpdateCropInspectorRequest struct {
	Name string `json:"name"`
}

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
		body := AddCropInspectorRequest{}
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, _ := queries.AddCropInspector(ctx, db.AddCropInspectorParams{body.Name, body.Usdaid})
		return c.JSON(res)
	})

	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		inspector, err := queries.GetCropInspector(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})

	cropinspector.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		name := UpdateCropInspectorRequest{}
		err = c.BodyParser(&name)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		inspector, err := queries.UpdateCropInspector(ctx, db.UpdateCropInspectorParams{
			Name:   name.Name,
			Usdaid: int32(id),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})

	// Get the crop inspectors for a district
	cropinspector.Get("/code/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		inspector, err := queries.GetInspectorForDistrict(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})

	// Delete a crop inspector by id
	cropinspector.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		inspector, err := queries.DeleteCropInspector(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(inspector)
	})
}
