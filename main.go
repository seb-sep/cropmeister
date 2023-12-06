package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/seb-sep/cropmeister/db"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := fiber.New()
	// ctx := context.Background()
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

	app.Get("/addValues", func(c *fiber.Ctx) error {
		//GenerateFakeInstances()// TODO: this is broken?!?!!??!
		return c.JSON("ARGH")
	})

	harvest := app.Group("/harvest")
	harvest.Get("", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Post("", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Get("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Put("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	harvest.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})

	crop := app.Group("/crop")
	crop.Get("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	crop.Put("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	crop.Delete("/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})

	farm := app.Group("/farm")
	farm.Get("", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Post("", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	farm.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	farm.Delete("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	purchase := app.Group("/purchase")
	purchase.Get("", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Post("", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	purchase.Delete("/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})

	app.Get("/cropinspector", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	app.Post("/cropinspector", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})

	farmer := app.Group("/farmer")
	farmer.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})
	farmer.Post("/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})

	districtcode := app.Group("/districtcode")
	districtcode.Get("", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Post("/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Get("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Put("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	districtcode.Delete("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	cropinspector := app.Group("/cropinspector")
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Get("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	cropinspector.Post("/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})

	app.Listen(":3000")
}
