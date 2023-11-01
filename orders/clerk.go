package orders

import (
	"backend/clerks"
	"backend/models"
	"backend/users"
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func clerk(r fiber.Router) {
  g := r.Group("/clerk")

  g.Get("/day", clerks.ClerkMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    orders, err := models.GetOrders(bson.M {
      "date": utils.GetToday(),
    }, bson.M {
      "time": -1,
    })

    if  err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(orders)
  })

  g.Post("/", clerks.ClerkMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
    pc := []models.ProductConfig{}
    json.Unmarshal(c.Body(), &pc)

    clerkID := fmt.Sprintf("%v", c.Locals("ID"))
    storeID := fmt.Sprintf("%v", c.Locals("storeID")) 

    ot := c.Query("type")
    ID := c.Query("ID")
    
    o := models.Order{}
    err := o.Create(ot, pc, storeID, clerkID, ID)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(o)
  })
}