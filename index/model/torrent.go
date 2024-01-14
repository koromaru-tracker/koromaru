package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Torrent struct {
	gorm.Model
	ID          uuid.UUID `gorm:"index:idx_torrent_id,unique"`
	Name        string    `form:"name" json:"name" binding:"required"`
	Description string    `form:"description" json:"description"`
}

func (t *Torrent) Create(db *gorm.DB) error {
	t.ID = uuid.New()
	return db.Create(t).Error
}
