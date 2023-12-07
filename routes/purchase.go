package routes

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddPurchaseRequest struct {
	CropType         string    `json:"crop_type"`
	FarmID           int32     `json:"farm_id"`
	PurchaseComplete bool      `json:"purchase_complete"`
	TotalQuantity    int32     `json:"total_quantity"`
	TotalPrice       float64   `json:"total_price"`
	PurchaseDate     time.Time `json:"purchase_date"`
}

func PurchaseRoutes(purchase fiber.Router) {
	ctx := context.Background()

	purchase.Get("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		purchases, err := queries.GetPurchases(ctx)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchases)
	})

	purchase.Post("", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		var body AddPurchaseRequest
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		res, err := queries.AddPurchase(ctx, db.AddPurchaseParams{
			CropType:         sql.NullString{String: body.CropType, Valid: true},
			FarmID:           sql.NullInt32{Int32: body.FarmID, Valid: true},
			PurchaseComplete: sql.NullBool{Bool: body.PurchaseComplete, Valid: true},
			TotalQuantity:    sql.NullInt32{Int32: body.TotalQuantity, Valid: true},
			TotalPrice:       sql.NullFloat64{Float64: body.TotalPrice, Valid: true},
			PurchaseDate:     sql.NullTime{Time: body.PurchaseDate, Valid: true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	purchase.Get("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		purchase, err := queries.GetPurchase(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchase)
	})

	purchase.Put("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		var body AddPurchaseRequest
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.UpdatePurchase(ctx, db.UpdatePurchaseParams{
			CropType:         sql.NullString{String: body.CropType, Valid: true},
			FarmID:           sql.NullInt32{Int32: body.FarmID, Valid: true},
			PurchaseComplete: sql.NullBool{Bool: body.PurchaseComplete, Valid: true},
			TotalQuantity:    sql.NullInt32{Int32: body.TotalQuantity, Valid: true},
			TotalPrice:       sql.NullFloat64{Float64: body.TotalPrice, Valid: true},
			PurchaseDate:     sql.NullTime{Time: body.PurchaseDate, Valid: true},
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)
	})

	purchase.Delete("/:id", func(c *fiber.Ctx) error {
		queries := c.Locals("db").(*db.Queries)
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		purchase, err := queries.DeletePurchase(ctx, int32(id))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(purchase)
	})

}
