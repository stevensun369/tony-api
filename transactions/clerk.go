package transactions

import "github.com/gofiber/fiber/v2"

func clerk(r fiber.Router) {
  g := r.Group("/clerk")

  g.Get("/test", func (c *fiber.Ctx) error {
    return c.JSON("hello")
  })
}