package links

import (
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetLink(c *fiber.Ctx) error {
	id := c.Params("id")
	var link models.Link
	if result := h.DB.First(&link, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&link)
}
