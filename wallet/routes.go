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

  g.Get("/product", func (c *fiber.Ctx) error {
    p := models.Product {
      ID: "0",
      Variants: map[string][]models.ProductVariant{
        // "size": {
        //   models.ProductVariant{Variant: "small"},
        //   models.ProductVariant{Variant: "medium"},
        // },
      },
      Options: map[string][]models.ProductOption{
        "flavor": {
        //   models.ProductOption{Option: "vanilla"},
        //   models.ProductOption{Option: "chocolate"},
        },
      },
    }

    p.AddVariant(p.ID, "size", models.ProductVariant{Variant: "small"})

    return c.JSON(p)
  })
}