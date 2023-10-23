package storeadmins

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func StoreAdminMiddleware(c *fiber.Ctx) error {
  token := c.Get("StoreAdminToken")

  if token == "" {
		return utils.MessageError(c, "no token")
	}

	storeAdmin := models.StoreAdmin{}
	if err := storeAdmin.ParseToken(token); err != nil {
		return utils.MessageError(c, err.Error())
	}

  c.Locals("storeID", storeAdmin.StoreID)

  return c.Next()
}

func Routes(r fiber.Router) {
  g := r.Group("/storeadmins")

  g.Get("/me", users.AuthMiddleware, StoreAdminMiddleware, func (c *fiber.Ctx) error {
    return c.JSON(fmt.Sprintf("%v", c.Locals("storeID")))
  })

  g.Post("/update", StoreAdminMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
		user := models.User {}
		storeAdmin := models.StoreAdmin {}
		ID := fmt.Sprintf("%v", c.Locals("ID"))

		// user
		err := user.Get(bson.M {
			"ID": ID, 
		})
		if err != nil {
			return utils.MessageError(c, err.Error())
		}
		token, err := user.GenToken()
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		// storeAdmin
		err = storeAdmin.Get(ID)
		if err != nil {
			return utils.MessageError(c, err.Error())
		}
		storeAdminToken, err := storeAdmin.GenToken()
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(bson.M{"token": token, "storeAdminToken": storeAdminToken, "user": user})
	})

  signup(g)
  login(g)
  store(g)
}