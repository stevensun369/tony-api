package orders

import (
	"backend/clerks"
	"backend/models"
	"backend/users"
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func clerk(r fiber.Router) {
  g := r.Group("/clerk")

  g.Post("/", clerks.ClerkMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    pc := []models.ProductConfig{}
    json.Unmarshal(c.Body(), &pc)

    clerkID := fmt.Sprintf("%v", c.Locals("ID"))
    storeID := fmt.Sprintf("%v", c.Locals("storeID")) 

    ot := c.Query("type")
    
    o := models.Order{}
    o.Build(pc)

    err := o.Create(ot, storeID, clerkID)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(pc)
  })
}