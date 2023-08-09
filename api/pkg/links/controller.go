package links

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := app.Group("/links")
	routes.Post("/", h.AddLink)
	routes.Get("/", h.GetLinks)
	routes.Get("/:id", h.GetLink)
	routes.Put("/:id", h.UpdateLink)
	routes.Delete("/:id", h.DeleteLink)
}