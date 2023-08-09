package links

import (
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetLinks(c *fiber.Ctx) error {
	var links []models.Link
	if result := h.DB.Find(&links); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&links)
}
