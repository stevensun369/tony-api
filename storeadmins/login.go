package storeadmins

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func login(r fiber.Router) {
	g := r.Group("/login")

	g.Post("/", users.AuthMiddleware, func(c *fiber.Ctx) error {
		// getting user ID
		ID := fmt.Sprintf("%v", c.Locals("ID"))

		// creating the storeAdmin
		sa := models.StoreAdmin{}
		sa.Get(ID)

		// getting storeAdmin token
		token, err := sa.GenToken()
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(bson.M{"storeAdminToken": token, "storeAdmin": sa})
	})
}
