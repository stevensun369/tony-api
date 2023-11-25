package main

import (
	"backend/clerks"
	"backend/db"
	"backend/env"
	"backend/orders"
	"backend/products"
	"backend/storeadmins"
	"backend/tests"
	"backend/users"
	"os"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	os.Mkdir("files", os.ModePerm)
	os.Mkdir("files/products", os.ModePerm)

	app := fiber.New(fiber.Config{
		Prefork:           !env.Dev,
		BodyLimit:         10 * 1024 * 1024,
		StreamRequestBody: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	db.InitDB(env.MongoURI)
	db.InitCache(env.RedisOptions)

	v := app.Group(fmt.Sprintf("%v", env.Version))

	v.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})

	tests.Routes(v)
	users.Routes(v)

	storeadmins.Routes(v)
	clerks.Routes(v)

	products.Routes(v)
	orders.Routes(v)

	app.Listen(":4200")
}
