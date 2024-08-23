package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&Ettevote{})
	db.AutoMigrate(&Aadress{})
	db.AutoMigrate(&Arinimi{})
	db.AutoMigrate(&Kapital{})
	db.AutoMigrate(&Majandusaasta{})
	db.AutoMigrate(&MarkusedKaardil{})
	db.AutoMigrate(&OiguslikVorm{})
	db.AutoMigrate(&Sidevahend{})
	db.AutoMigrate(&Staatus{})
	db.AutoMigrate(&TeatatudTegevusala{})
	db.AutoMigrate(&Pohikiri{})

	return db, nil
}

func InsertBatch[T any](db *gorm.DB, items *[]T, batchSize int) {
	if len(*items) >= batchSize {
		db.Create(items)
		*items = (*items)[:0]
	}
}
func InsertAll[T any](db *gorm.DB, items *[]T) {
	if len(*items) > 0 {
		db.Create(items)
		*items = (*items)[:0]
	}
}
