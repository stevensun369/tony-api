package wallet

import (
	"backend/models"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router) {
  g := r.Group("/wallet")

  g.Get("/test", func (c *fiber.Ctx) error {
    wallet := models.Wallet{}
    wallet.ID = "fUL0V7XBj8"

    // wallet.In(30)
    err := wallet.Out(10)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }
    return c.JSON(wallet)
  })
}