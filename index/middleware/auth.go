package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusUnauthorized)
			return c.Redirect("/login")
		},
		SigningKey: []byte(secret),
	})
}
