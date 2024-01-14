package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Configure(app *fiber.App) *fiber.App {

	// Add logger middleware
	app.Use(logger.New())

	// Add monitor middleware
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Koromaru Tracker Metrics"}))

	return app
}
