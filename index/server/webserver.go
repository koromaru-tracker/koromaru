package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/api"
	"github.com/koromaru-tracker/koromaru/index/types"
	"gorm.io/gorm"
)

func Serve(cfg *types.Config, db *gorm.DB) {
	app := fiber.New()

	// Set Database Connection to Fiber Context
	app.Use(func(c *fiber.Ctx) error {
		index.SetLocal[*gorm.DB](c, "db", db)
		// Go to next middleware:
		return c.Next()
	})

	app = api.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Webserver.Port))
}
