package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
)

type AddPurchaseRequest struct {
	CropType         string  `json:"cropType"`
	FarmID           int32   `json:"farmId"`
	FarmerName       string  `json:"farmerName"`
	PurchaseComplete bool    `json:"purchaseComplete"`
	TotalQuantity    int32   `json:"totalQuantity"`
	TotalPrice       float64 `json:"totalPrice"`
	PurchaseDate     YMD     `json:"purchaseDate"`
}

type YMD struct {
	Date time.Time
}

func (y *YMD) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	y.Date = t
	return nil
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
		body := AddPurchaseRequest{}
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		res, err := queries.AddPurchase(ctx, db.AddPurchaseParams{
			CropType:         sql.NullString{String: body.CropType, Valid: true},
			FarmID:           sql.NullInt32{Int32: body.FarmID, Valid: true},
			FarmerName:       body.FarmerName,
			PurchaseComplete: sql.NullBool{Bool: body.PurchaseComplete, Valid: true},
			TotalQuantity:    sql.NullInt32{Int32: body.TotalQuantity, Valid: true},
			TotalPrice:       sql.NullFloat64{Float64: body.TotalPrice, Valid: true},
			PurchaseDate:     sql.NullTime{Time: body.PurchaseDate.Date, Valid: true},
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
		body := AddPurchaseRequest{}
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		res, err := queries.UpdatePurchase(ctx, db.UpdatePurchaseParams{
			CropType:         sql.NullString{String: body.CropType, Valid: true},
			FarmID:           sql.NullInt32{Int32: body.FarmID, Valid: true},
			PurchaseComplete: sql.NullBool{Bool: body.PurchaseComplete, Valid: true},
			TotalQuantity:    sql.NullInt32{Int32: body.TotalQuantity, Valid: true},
			TotalPrice:       sql.NullFloat64{Float64: body.TotalPrice, Valid: true},
			PurchaseDate:     sql.NullTime{Time: body.PurchaseDate.Date, Valid: true},
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
