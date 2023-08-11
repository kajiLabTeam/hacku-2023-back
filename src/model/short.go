package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Short struct {
	ID               int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	UserID           string `gorm:"type:varchar(28)" json:"userId"`
	GenreID          int    `json:"genreId"`
	Title            string
	CreatedAt        time.Time
	BrowsingHistorys []BrowsingHistory `gorm:"foreignkey:ShoatID"`
	Slides           []Slide           `gorm:"foreignkey:ShoatID"`
	Tags             []Tag             `gorm:"foreignkey:ShoatID"`
}

func GetShoatByID(id int) *Short {
	shoat := Short{}
	result := db.First(&shoat, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &shoat
}

func InsertShoat(short Short) {
	db.Create(&short)
}

func DeleatShoat(id int) {
	short := Short{}
	db.Delete(&short, id)
}
