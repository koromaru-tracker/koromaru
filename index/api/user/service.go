package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	user := new(model.User)
	user.Username = "test"
	user.Roles = []model.Role{
		{
			ID: uuid.Must(uuid.Parse("a87cc8f5-81e6-4702-a9c5-a80a5faf7ec4")),
		},
	}
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

func GetPassKeys(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")

	user := new(model.User)
	users, err := user.GetAll(db)
	if err != nil {
		return err
	}

	var passkeys []string
	for _, u := range users {
		passkeys = append(passkeys, u.PassKey)
	}

	return c.JSON(passkeys)
}
