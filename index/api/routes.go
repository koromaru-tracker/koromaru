package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index/api/torrent"
	"github.com/koromaru-tracker/koromaru/index/api/user"
)

func RegisterRoutes(app *fiber.App) *fiber.App {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User Routes
	v1.Get("/users", user.Get)
	v1.Post("/users", user.Create)

	// Torrent Routes
	v1.Get("/torrents", torrent.Get)

	return app
}
