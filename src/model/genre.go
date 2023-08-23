package model

import (
	"errors"

	"gorm.io/gorm"
)

type Genre struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	GenreName string `json:"genreName"`
	Short     []Short  `gorm:"foreignkey:GenreID"`
}

func GetAllGenre() []Genre {
	g := []Genre{}
	result := db.Find(&g)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return g
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

func DeleteGenre(id int) {
	g := Genre{}
	db.Delete(&g, id)
}
