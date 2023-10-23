package tests

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router) {
  g := r.Group("/tests")

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