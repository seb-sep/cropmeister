package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"
	"github.com/seb-sep/cropmeister/routes"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbUrl = os.Getenv("DATABASE_URL")

func main() {

	// app init
	app := fiber.New()
	dbUrl := os.Getenv("DATABASE_URL")
	msql, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer msql.Close()
	queries := db.New(msql)

	// add the db to the fiber context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", queries)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

	// routing groups
	harvest := app.Group("/harvest")
	routes.HarvestRoutes(harvest)

	crop := app.Group("/crop")
	routes.CropRoutes(crop)

	farm := app.Group("/farm")
	routes.FarmRoutes(farm)

	purchase := app.Group("/purchase")
	routes.PurchaseRoutes(purchase)

	farmer := app.Group("/farmer")
	routes.FarmerRoutes(farmer)

	districtcode := app.Group("/districtcode")
	routes.DistrictCodeRoutes(districtcode)

	cropinspector := app.Group("/cropinspector")
	routes.CropInspectorRoutes(cropinspector)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}
}
