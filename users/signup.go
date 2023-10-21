package users

import (
	"github.com/gofiber/fiber/v2"
)

func signup(r fiber.Router) {
  r.Post("/signup/phone", func (c *fiber.Ctx) error {
    return c.JSON("testing")
  })

  
}