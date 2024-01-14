package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	user := new(model.User)
	user.Username = "test"
	user.Password = "test"
	user.Email = "test@test.de"

	err := user.Create(db)
	if err != nil {
		return err
	}

	return c.SendString("User created 👋!")
}

func Get(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	user := new(model.User)
	users, err := user.GetAll(db)
	if err != nil {
		return err
	}

	return c.JSON(users)
}
