package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Shoat struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	UserID    []Keyword `gorm:"foreignkey:ID"`
	GenreID   []Genre   `gorm:"foreignkey:ID"`
	Title     string
	CreatedAt time.Time
}

func GetShoatByID(id int) *Shoat {
	shoat := Shoat{}
	result := db.First(&shoat, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &shoat
}

func InsertShoat(shoat Shoat) {
	db.Create(&shoat)
}

func DeleatShoat(id int) {
	shoat := Shoat{}
	db.Delete(&shoat, id)
}
