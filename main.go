package main

import (
	"backend/db"
	"backend/env"
	"backend/users"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	db.InitDB(env.MongoURI)
	db.InitCache(env.RedisOptions)

	app.Get("/v0/test", func(c *fiber.Ctx) error {
		return c.JSON("ok!")
	})

	users.Routes(app)

	app.Listen(":9000")
}
