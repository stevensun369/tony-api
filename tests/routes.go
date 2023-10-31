package tests

import (
	"backend/models"
	"backend/utils"
	"fmt"

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

  g.Post("/image", func (c *fiber.Ctx) error {
    fmt.Println("in")
    file, err := c.FormFile("image")
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    c.SaveFile(file, fmt.Sprintf("./files/products/%s", file.Filename))

    return c.JSON(file.Filename)
  })

  g.Get("/image", func (c *fiber.Ctx) error {
    image := c.Query("image")
    return c.JSON(image)
  })
}