package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/seb-sep/cropmeister/db"
)

func main() {
	app := fiber.New()
	ctx := context.Background()
	dbUrl := os.Getenv("DATABASE_URL")
	msql, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer msql.Close()

	queries := db.New(msql)

	app.Get("/", func(c *fiber.Ctx) error {

		authors, _ := queries.ListAuthors(ctx)
		names := []string{}
		for _, author := range authors {
			names = append(names, author.Name)
		}
		return c.JSON(names)
	})

	app.Get("/harvest", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})

	app.Post("/harvest", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})

	app.Get("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	app.Put("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})
	app.Delete("/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("crop")
	})

	app.Get("/farm", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})
	app.Post("/farm", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	app.Get("/farm/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	app.Put("/farm/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	app.Delete("/farm/:id", func(c *fiber.Ctx) error {
		return c.JSON("farm")
	})

	app.Get("/harvest/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	app.Put("/harvest/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})
	app.Delete("/harvest/:type", func(c *fiber.Ctx) error {
		return c.JSON("harvest")
	})

	app.Get("/purchase", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})

	app.Post("/purchase", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})

	app.Get("/purchase/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	app.Put("/purchase/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})
	app.Delete("/purchase/:id", func(c *fiber.Ctx) error {
		return c.JSON("purchase")
	})

	app.Get("/farmer/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})

	app.Post("/farmer/:id", func(c *fiber.Ctx) error {
		return c.JSON("farmer")
	})

	app.Get("/cropinspector", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	app.Post("/cropinspector", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})

	app.Get("/districtcode/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	app.Post("/districtcode/crop/:type", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	app.Get("/districtcode/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	app.Post("/districtcode/inspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	app.Get("/districtcode/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	app.Put("/districtcode/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})
	app.Delete("/districtcode/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	app.Get("/districtcode", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	app.Post("/districtcode", func(c *fiber.Ctx) error {
		return c.JSON("districtcode")
	})

	app.Get("/cropinspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	app.Put("/cropinspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	app.Get("/cropinspector/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})

	app.Get("/cropinspector/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})
	app.Post("/cropinspector/code/:id", func(c *fiber.Ctx) error {
		return c.JSON("cropinspector")
	})

	app.Listen(":3000")
}
