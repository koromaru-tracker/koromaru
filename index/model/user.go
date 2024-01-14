package model

import (
	"github.com/google/uuid"
	"github.com/koromaru-tracker/koromaru/index/security"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey,index:idx_user_id,unique"`
	Username string    `gorm:"unique"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
	Password string
	Email    string `gorm:"unique"`
	PassKey  string `gorm:"index:idx_user_passkey"`
}

func (u *User) Create(db *gorm.DB) error {
	u.ID = uuid.New()

	// hash username + salt and save as passkey
	passkey, err := security.GeneratePassKeyHash(u.Username)
	if err != nil {
		return err
	}

	// hash password
	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	u.PassKey = passkey
	return db.Create(u).Error
}

func (u *User) GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Roles").Find(&users).Error
	return users, err
}
