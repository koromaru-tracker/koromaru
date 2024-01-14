package page

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"github.com/koromaru-tracker/koromaru/index/types"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	}, "layouts/enter")
}

func RegisterRequest(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	// Extract the credentials from the request body
	registerRequest := new(types.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if user exists
	var user model.User
	db.First(&user, "username = ?", registerRequest.Username)
	if user.Username != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// Check if email exists
	db.First(&user, "email = ?", registerRequest.Email)
	if user.Email != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	// Check if passwords match
	fmt.Println(registerRequest.Password)
	fmt.Println(registerRequest.PasswordConfirm)
	if registerRequest.Password != registerRequest.PasswordConfirm {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Passwords do not match",
		})
	}

	// Set default roles if not first user
	var roles []model.Role

	// count users
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		roles = []model.Role{
			{
				ID: uuid.Must(uuid.Parse("a87cc8f5-81e6-4702-a9c5-a80a5faf7ec4")),
			},
		}
	} else {
		roles = []model.Role{
			{
				ID: uuid.Must(uuid.Parse("17ef6f73-cfc7-432e-b20b-7569b53d932d")),
			},
		}
	}

	// Create user
	user = model.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
		Email:    registerRequest.Email,
		Roles:    roles,
	}

	user.Create(db)

	return c.Redirect("/login")
}
