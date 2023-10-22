package main

import (
	"backend/db"
	"backend/env"
	"backend/storeadmin"
	"backend/users"
	"backend/wallet"
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

	v := app.Group(fmt.Sprintf("%v", env.Version))

	v.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON("ok!")
	})

	users.Routes(v)
	wallet.Routes(v)
	storeadmin.Routes(v)

	app.Listen(":9000")
}
