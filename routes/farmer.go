package routes

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddFarmerRequest struct {
	Name     string  `json:"name"`
	Budget   float64 `json:"budget"`
	NetWorth float64 `json:"netWorth"`
	FarmID   int32   `json:"farmId"`
}

func FarmerRoutes(farmer fiber.Router) {
	ctx := context.Background()

	farmer.Get("/:name/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		name := c.Params("name")
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		farmer, err := queries.GetFarmer(ctx, db.GetFarmerParams{FarmID: int32(id), Name: name})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farmer)
	})

	farmer.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		body := AddFarmerRequest{}
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, _ := queries.AddFarmer(ctx, db.AddFarmerParams{
			Name:     body.Name,
			Budget:   sql.NullFloat64{body.Budget, true},
			NetWorth: sql.NullFloat64{body.NetWorth, true},
			FarmID:   body.FarmID,
		})
		return c.JSON(res)
	})

	farmer.Delete("/:name", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		farmers, err := queries.DeleteFarmers(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farmers)
	})

}
