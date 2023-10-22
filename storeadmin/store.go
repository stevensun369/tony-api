package storeadmin

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func store(r fiber.Router) {
  g := r.Group("/store")

  g.Get("/", users.AuthMiddleware, StoreAdminMiddleware, func (c *fiber.Ctx) error {
    storeID := fmt.Sprintf("%v", c.Locals("storeID"))

    s := models.Store{}
    s.Get(storeID)

    return c.JSON(s)
  })

  g.Post("/name", users.AuthMiddleware, StoreAdminMiddleware, func (c *fiber.Ctx) error {
    var body map[string]string
    json.Unmarshal(c.Body(), &body)

    storeID := fmt.Sprintf("%v", c.Locals("storeID"))

    s := models.Store{}

    if err := s.ChangeName(storeID, body["name"]); err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(s)
  })

  g.Post("/open", users.AuthMiddleware, StoreAdminMiddleware, func (c *fiber.Ctx) error {
    var body map[string]string
    json.Unmarshal(c.Body(), &body)

    storeID := fmt.Sprintf("%v", c.Locals("storeID"))

    s := models.Store{}

    if err := s.ChangeOpen(storeID, body["open"]); err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(s)
  })
}