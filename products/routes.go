package products

import "github.com/gofiber/fiber/v2"

func Routes(r fiber.Router) {
	g := r.Group("/products")

	change(g)
	get(g)
}
