package backend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"gorm.io/gorm"
)

func UserList(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")
	var user model.User
	var users []model.User

	users, err := user.GetAll(db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Render("backend/user/list", fiber.Map{
		"Title": "List User",
		"Users": users,
	})
}

func UserLogout(c *fiber.Ctx) error {
	// Remove Cookie Token
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	c.Cookie(cookie)

	return c.Redirect("/login")
}
