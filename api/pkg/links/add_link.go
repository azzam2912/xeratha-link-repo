package links

import (
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddLinkRequestBody struct {
	Title       string `json:"title"`
	LinkAddress string `json:"linkAddress"`
	Tags        string `json:"tags"`
}

func (h handler) AddLink(c *fiber.Ctx) error {
	body := AddLinkRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var link models.Link
	link.Title = body.Title
	link.LinkAddress = body.LinkAddress
	link.Tags = body.Tags
	if result := h.DB.Create(&link); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(&link)
}
