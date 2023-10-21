package users

import (
	"backend/env"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
  g := app.Group(fmt.Sprintf("/%v/users", env.Version))

  signup(g)
}