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
	NetWorth float64 `json:"net_worth"`
	FarmID   int32   `json:"farm_id"`
}

func FarmerRoutes(farmer fiber.Router) {
	ctx := context.Background()

	farmer.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		var body map[string]string
		err = c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		farmers, err := queries.GetFarmer(ctx, db.GetFarmerParams{FarmID: int32(id), Name: body["name"]})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(farmers)
	})

	farmer.Post("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		var body AddFarmerRequest
		err := c.BodyParser(body)
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

}
