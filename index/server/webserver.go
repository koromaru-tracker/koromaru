package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/api"
	"github.com/koromaru-tracker/koromaru/index/frontend"
	"github.com/koromaru-tracker/koromaru/index/types"
	"gorm.io/gorm"
)

func Serve(cfg *types.Config, db *gorm.DB) {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		// Views Layout is the global layout for all template render until override on Render function.
		ViewsLayout: "layouts/main",
	})

	// Set Database Connection to Fiber Context
	app.Use(func(c *fiber.Ctx) error {
		index.SetLocal[*gorm.DB](c, "db", db)
		// Go to next middleware:
		return c.Next()
	})

	app.Static("/", "./dist")

	// Configure Fiber App
	app = Configure(app)

	app = api.RegisterRoutes(app)
	app = frontend.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Webserver.Port))
}
