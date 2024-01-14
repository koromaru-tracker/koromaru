package frontend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index/frontend/backend"
	"github.com/koromaru-tracker/koromaru/index/frontend/page"
	"github.com/koromaru-tracker/koromaru/index/middleware"
)

func RegisterRoutes(app *fiber.App) *fiber.App {
	app.Get("/", page.Index)
	app.Get("/login", page.Login)
	app.Post("/login", page.LoginRequest)
	app.Get("/register", page.Register)
	app.Post("/register", page.RegisterRequest)

	jwt := middleware.NewAuthMiddleware("test")

	backendGroup := app.Group("/backend")
	backendGroup.Get("/", jwt, backend.Index)

	torrentGroup := backendGroup.Group("/torrent")
	torrentGroup.Get("/list", jwt, backend.TorrentList)
	torrentGroup.Get("/create", jwt, backend.TorrentCreateView)
	torrentGroup.Post("/create", jwt, backend.TorrentCreate)
	torrentGroup.Get("/view/:id", jwt, backend.TorrentView)

	userGroup := backendGroup.Group("/user")
	userGroup.Get("/list", jwt, backend.UserList)
	userGroup.Get("/logout", jwt, backend.UserLogout)

	return app
}
