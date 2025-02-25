package products

import (
	"backend/models"
	"backend/storeadmins"
	"backend/users"
	"backend/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func options(r fiber.Router) {
	g := r.Group("/options")

	g.Post("/", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Query("key")

		p := models.Product{}
		if err := p.CreateOptionKey(productID, key); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Delete("/", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Query("key")

		p := models.Product{}
		if err := p.RemoveOptionKey(productID, key); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Post("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")

		option := models.ProductOption{}
		json.Unmarshal(c.Body(), &option)

		p := models.Product{}
		if err := p.AddOption(productID, key, option); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Delete("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")
		option := c.Query("option")

		p := models.Product{}
		if err := p.RemoveOption(productID, key, option); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Patch("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")

		option := models.ProductOption{}
		json.Unmarshal(c.Body(), &option)

		p := models.Product{}
		if err := p.ChangeOption(productID, key, option); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})
}
