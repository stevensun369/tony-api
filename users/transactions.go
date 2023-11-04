package users

import (
	"backend/models"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func transactions(r fiber.Router) {
  g := r.Group("/transactions")

  g.Get("/", AuthMiddleware, func (c *fiber.Ctx) error {
    user := models.User {}
    utils.GetLocals(c, "user", &user)
    
    transactions, err := models.GetTransactions(bson.M {
      "to": user.WalletID,
    }, bson.D{
      {Key: "date", Value: -1}, 
      {Key: "time", Value: -1},
    })

    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(transactions)
  })
}