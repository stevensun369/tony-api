package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
  fmt.Println("hello world")

  app := fiber.New()

  app.Get("/test", func (c *fiber.Ctx) error {
    return c.JSON("ok!")
  })
}