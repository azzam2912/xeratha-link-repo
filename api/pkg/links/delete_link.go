package links

import (
    "github.com/gofiber/fiber/v2"
    "github.com/azzam2912/xeratha-link-repo/api/pkg/common/models"
)

func (h handler) DeleteLink(c *fiber.Ctx) error {
	id := c.Params("id")
	var link models.Link
	if result := h.DB.First(&link, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Delete(&link)
	return c.SendStatus(fiber.StatusOK)
}