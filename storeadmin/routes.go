package storeadmin

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StoreAdminMiddleware(c *fiber.Ctx) error {
  token := c.Get("StoreAdminToken")

  if token == "" {
		return utils.MessageError(c, "no token")
	}

	storeAdmin := models.StoreAdmin{}
	if err := storeAdmin.ParseToken(token); err != nil {
		return utils.MessageError(c, err.Error())
	}

  c.Locals("storeID", storeAdmin.StoreID)

  return c.Next()
}

func Routes(r fiber.Router) {
  g := r.Group("/storeadmin")

  g.Get("/me", users.AuthMiddleware, StoreAdminMiddleware, func (c *fiber.Ctx) error {
    return c.JSON(fmt.Sprintf("%v", c.Locals("storeID")))
  })

  signup(g)
  login(g)
  store(g)
}