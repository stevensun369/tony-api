package users

import (
	"backend/env"
	"backend/models"
	"backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
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
	
	g := r.Group(fmt.Sprintf("/%v/users", env.Version))
	
	g.Get("/me", AuthMiddleware, func (c *fiber.Ctx) error {
		user := models.User {}

		utils.GetLocals(c, "user", &user)
		return c.JSON(user)
	})
	
  signup(g)
	login(g)
}