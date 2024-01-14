package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/koromaru-tracker/koromaru/index/types"
)

func Connect(cfg *types.Config) (*gorm.DB, error) {
	if cfg.Database.Provider == "sqlite" {
		return newSqliteDB(cfg.Database.Path)
	}

	// return Error if no database provider is found
	return nil, fmt.Errorf("no database provider found")
}

func newSqliteDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
