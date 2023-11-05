package users

import (
	"backend/env"
	"backend/models"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	
	if token == "" {
		return utils.MessageError(c, "no token")
	}

	user := models.User{}
	if err := user.ParseToken(token); err != nil {
		return utils.MessageError(c, err.Error())
	}

	c.Locals("ID", user.ID)
	utils.SetLocals(c, "user", user)

	return c.Next()
}

func Routes(r fiber.Router) {
	
	g := r.Group("/users")
	
	g.Get("/me", AuthMiddleware, func (c *fiber.Ctx) error {
		user := models.User {}

		utils.GetLocals(c, "user", &user)
		return c.JSON(user)
	})

	g.Post("/update", AuthMiddleware, func (c *fiber.Ctx) error {
		user := models.User {}
		ID := fmt.Sprintf("%v", c.Locals("ID"))

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

		return c.JSON(bson.M{"token": token})
	})
	
	g.Get("/check", func (c *fiber.Ctx) error {
		phone := c.Query("phone")
		// implenting the demo features
		// for play store and appstore approval
		if phone == env.DemoPhone {
			return c.JSON(
				bson.M {
					"check": true,
				},
			)
		}

		check := models.UserCheck(
			bson.M {
				"phone": phone,
			},
		)

		return c.JSON(
			bson.M {
				"check": check,
			},
		)
	})

	g.Get("/wallet", AuthMiddleware, func (c *fiber.Ctx) error {
    user := models.User {}
    utils.GetLocals(c, "user", &user)

    wallet := models.Wallet {}
    err := wallet.Get(user.WalletID)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(wallet)
  })

  signup(g)
	login(g)
	transactions(g)
}