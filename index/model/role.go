package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primaryKey,index:idx_role_id,unique"`
	Name string    `gorm:"unique"`
}
