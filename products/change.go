package products

import (
	"backend/models"
	"backend/storeadmins"
	"backend/users"
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func change(r fiber.Router) {
  g := r.Group("/")

  g.Post("/", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    p := models.Product{}
    json.Unmarshal(c.Body(), &p)

    storeID := fmt.Sprintf("%v", c.Locals("storeID"))
    p.StoreID = storeID

    if err := p.Create(); err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(p)
  })

  g.Put("/tags", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    productID := c.Query("productID")
    tag := c.Query("tag")

    p := models.Product{}
    err := p.AddTag(productID, tag)

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(p)
  })

  g.Delete("/tags", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
    productID := c.Query("productID")
    tag := c.Query("tag")

    p := models.Product{}
    err := p.RemoveTag(productID, tag)

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(p)
  })

  g.Put("/:field", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    var body map[string]interface{}
    json.Unmarshal(c.Body(), &body)

    field := c.Params("field")
    productID := c.Query("productID")

    p := models.Product{}

    err := p.UpdateField(productID, field, body[field])

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON("ok")
  })

  variants(g)
  options(g)
}