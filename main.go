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

	app.Listen(":3000")
}
