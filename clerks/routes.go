package clerks

import (
	"backend/models"
	"backend/users"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ClerkMiddleware(c *fiber.Ctx) error {
  token := c.Get("ClerkToken")

  if token == "" {
		return utils.MessageError(c, "no token")
	}

	clerk := models.Clerk{}
	if err := clerk.ParseToken(token); err != nil {
		return utils.MessageError(c, err.Error())
	}

  c.Locals("storeID", clerk.StoreID)

  return c.Next()
}

func Routes(r fiber.Router) {
  g := r.Group("/clerks")

  g.Get("/me", users.AuthMiddleware, ClerkMiddleware, func (c *fiber.Ctx) error {
    return c.JSON(fmt.Sprintf("%v", c.Locals("storeID")))
  })

	g.Post("/update", ClerkMiddleware, users.AuthMiddleware, func (c *fiber.Ctx) error {
		user := models.User {}
		clerk := models.Clerk {}
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

		// clerk
		err = clerk.Get(ID)
		if err != nil {
			return utils.MessageError(c, err.Error())
		}
		clerkToken, err := clerk.GenToken()
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(bson.M{"token": token, "clerkToken": clerkToken})
	})

	g.Get("/store", ClerkMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		storeID := fmt.Sprintf("%v", c.Locals("storeID"))
		
		store := models.Store{}
		err := store.Get(storeID)
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(store)
	})

  signup(g)
  login(g)
}