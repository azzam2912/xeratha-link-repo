package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/config"
	"github.com/azzam2912/xeratha-link-repo/api/pkg/common/db"
	"github.com/azzam2912/xeratha-link-repo/api/pkg/links"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Config failed", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	links.RegisterRoutes(app, db)
	
	app.Listen(c.Port)
}
