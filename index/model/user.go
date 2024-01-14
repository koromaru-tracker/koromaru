package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"index:idx_user_id,unique"`
	Username string
	Password string
	Email    string
	PassKey  string
}
