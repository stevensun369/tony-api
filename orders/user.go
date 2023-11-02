package orders

import (
	"backend/models"
	"backend/users"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func user(r fiber.Router) {
  g := r.Group("/user")

  g.Get("/", users.AuthMiddleware, func (c *fiber.Ctx) error {
    orderID := c.Query("orderID")

    order := models.Order {}
    err := order.GetOrder(orderID)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(order)
  })
}