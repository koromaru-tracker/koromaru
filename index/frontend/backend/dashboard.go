package backend

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Index(c *fiber.Ctx) error {
	// Check if user is authenticated
	user := c.Locals("user").(*jtoken.Token)
	if user == nil {
		return c.Redirect("/login")
	}
	claims := user.Claims.(jtoken.MapClaims)
	if claims["role"] != "admin" {
		fmt.Println("Not admin")
	}

	return c.Render("backend/dashboard", fiber.Map{
		"Title": "Dasboard",
	})
}
