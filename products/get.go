package products

import (
	"backend/models"
	"backend/users"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func get(r fiber.Router) {
  g := r.Group("/")

  g.Get("/product", users.AuthMiddleware, func (c *fiber.Ctx) error {
    productID := c.Query("productID")

    p := models.Product {}
    err := p.Get(productID)

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(p)
  })

  g.Get("/tag", users.AuthMiddleware, func (c *fiber.Ctx) error {
    tag := c.Query("tag")
    storeID := c.Query("storeID")

    products, err := models.GetProducts(bson.M {
      "tags": tag,
      "storeID": storeID,
    })

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(products)
  })

  g.Get("/", users.AuthMiddleware, func (c *fiber.Ctx) error {
    storeID := c.Query("storeID")

    products, err := models.GetProducts(bson.M {
      "storeID": storeID,
    })

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(products)
  })
}