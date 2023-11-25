package products

import (
	"backend/models"
	"backend/storeadmins"
	"backend/users"
	"backend/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func variants(r fiber.Router) {
	g := r.Group("/variants")

	g.Post("/", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Query("key")

		p := models.Product{}
		if err := p.CreateVariantKey(productID, key); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Delete("/", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Query("key")

		p := models.Product{}
		if err := p.RemoveVariantKey(productID, key); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Post("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")

		variant := models.ProductVariant{}
		json.Unmarshal(c.Body(), &variant)

		p := models.Product{}
		if err := p.AddVariant(productID, key, variant); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Delete("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")
		variant := c.Query("variant")

		p := models.Product{}
		if err := p.RemoveVariant(productID, key, variant); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})

	g.Patch("/:key", storeadmins.StoreAdminMiddleware, users.AuthMiddleware, func(c *fiber.Ctx) error {
		productID := c.Query("productID")
		key := c.Params("key")

		variant := models.ProductVariant{}
		json.Unmarshal(c.Body(), &variant)

		p := models.Product{}
		if err := p.ChangeVariant(productID, key, variant); err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(p)
	})
}
