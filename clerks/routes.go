package clerks

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ClerkMiddleware(c *fiber.Ctx) error {
  token := c.Get("ClerkToken")

  if token == "" {
		return utils.MessageError(c, "no token")
	}

	clerk := models.Clerk{}
	if err := clerk.ParseToken(token); err != nil {
		return utils.MessageError(c, err.Error())
	}

  c.Locals("storeID", clerk.StoreID)

  return c.Next()
}

func Routes(r fiber.Router) {
  g := r.Group("/clerks")

  g.Get("/me", users.AuthMiddleware, ClerkMiddleware, func (c *fiber.Ctx) error {
    return c.JSON(fmt.Sprintf("%v", c.Locals("storeID")))
  })

  signup(g)
  login(g)
}