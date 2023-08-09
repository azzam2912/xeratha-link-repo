package links

import (
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateLinkRequestBody struct {
	Title			string 	`json:"title"`
	LinkAddress		string	`json:"linkAddress"`
	Tags			string 	`json:"tags"`
}

func (h handler) UpdateLink(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateLinkRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var link models.Link
	if result := h.DB.First(&link, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	link.Title = body.Title
	link.LinkAddress = body.LinkAddress
	link.Tags = body.Tags
	h.DB.Save(&link)
	return c.Status(fiber.StatusOK).JSON(&link)
}