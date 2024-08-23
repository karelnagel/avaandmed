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

	db.AutoMigrate(&Yldandmed{})
	db.AutoMigrate(&Aadress{})
	db.AutoMigrate(&Arinimi{})
	db.AutoMigrate(&Kapital{})
	db.AutoMigrate(&Majandusaasta{})
	db.AutoMigrate(&MarkusKaardil{})
	db.AutoMigrate(&OiguslikVorm{})
	db.AutoMigrate(&Sidevahend{})
	db.AutoMigrate(&Staatus{})
	db.AutoMigrate(&TeatatudTegevusala{})
	db.AutoMigrate(&Pohikiri{})
	db.AutoMigrate(&InfoMajandusaastaAruandest{})

	db.AutoMigrate(&KaardileKantudIsik{})

	db.AutoMigrate(&KandevalineIsik{})

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

type Tabler interface {
	TableName() string
}

func (Yldandmed) TableName() string {
	return "ettevotted"
}
func (Aadress) TableName() string {
	return "aadressid"
}
func (Arinimi) TableName() string {
	return "arinimed"
}
func (Kapital) TableName() string {
	return "kapitalid"
}
func (Majandusaasta) TableName() string {
	return "majandusaastad"
}
func (MarkusKaardil) TableName() string {
	return "markused_kaardil"
}
func (Staatus) TableName() string {
	return "staatused"
}
func (Sidevahend) TableName() string {
	return "sidevahendid"
}
func (InfoMajandusaastaAruandest) TableName() string {
	return "info_majandusaasta_aruandest"
}
func (TeatatudTegevusala) TableName() string {
	return "teatatud_tegevusalad"
}
func (Pohikiri) TableName() string {
	return "pohikirjad"
}
func (OiguslikVorm) TableName() string {
	return "oiguslikud_vormid"
}
func (KaardileKantudIsik) TableName() string {
	return "kaardile_kantud_isikud"
}
func (KandevalineIsik) TableName() string {
	return "kandevalised_isikud"
}
