package orders

import "github.com/gofiber/fiber/v2"

func Routes(r fiber.Router) {
	g := r.Group("/orders")

	clerk(g)
	user(g)
}
