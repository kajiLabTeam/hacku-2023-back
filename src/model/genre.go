package model

import (
	"errors"

	"gorm.io/gorm"
)

type Genre struct {
	ID        int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	GenreName string `json:"genreName"`
	Shoat     Short  `gorm:"foreignkey:GenreID"`
}

func GetGenreByID(id int) *Genre {
	genre := Genre{}
	result := db.First(&genre, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &genre
}

func InsertGenre(genre Genre) {
	db.Create(&genre)
}

func DeleatGenret(id int) {
	genre := Genre{}
	db.Delete(&genre, id)
}
