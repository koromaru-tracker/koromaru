package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index/types"
	"gorm.io/gorm"
)

func Serve(cfg *types.Config, db *gorm.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + cfg.Webserver.Port))
}
