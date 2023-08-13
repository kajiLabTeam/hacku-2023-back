package model

import (
	"errors"

	"gorm.io/gorm"
)

type Genre struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	GenreName string `json:"genreName"`
	Shoat     Short  `gorm:"foreignkey:GenreID"`
}

func GetGenreByID(id int) *Genre {
	g := Genre{}
	result := db.First(&g, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &g
}

func InsertGenre(g Genre) {
	db.Create(&g)
}

func DeleatGenret(id int) {
	g := Genre{}
	db.Delete(&g, id)
}
