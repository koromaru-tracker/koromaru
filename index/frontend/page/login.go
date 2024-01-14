package page

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"github.com/koromaru-tracker/koromaru/index/security"
	"github.com/koromaru-tracker/koromaru/index/types"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func LoginRequest(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	// Extract the credentials from the request body
	loginRequest := new(types.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find the user by credentials
	var user model.User
	db.First(&user, "username = ?", loginRequest.Username)
	if user.Username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Check if password is correct
	passwordIsCorrect := security.CheckPasswordHash(loginRequest.Password, user.Password)
	if !passwordIsCorrect {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":      user.ID,
		"user":    user.Username,
		"passkey": user.PassKey,
		"roles":   user.Roles,
		"exp":     time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("test"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	tookenCookie := new(fiber.Cookie)
	tookenCookie.Name = "token"
	tookenCookie.Value = t

	c.Cookie(tookenCookie)
	return c.Redirect("/backend")
}
