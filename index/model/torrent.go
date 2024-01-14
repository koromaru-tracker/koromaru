package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Torrent struct {
	gorm.Model
	ID          uuid.UUID `gorm:"index:idx_torrent_id,unique"`
	Name        string
	Description string
}
