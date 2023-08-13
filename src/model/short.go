package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Short struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           string `gorm:"type:varchar(28)" json:"userId"`
	GenreID          int    `json:"genreId"`
	Title            string
	CreatedAt        time.Time
	BrowsingHistories []BrowsingHistory `gorm:"foreignkey:ShoatID"`
	Slides           []Slide           `gorm:"foreignkey:ShoatID"`
	Tags             []Tag             `gorm:"foreignkey:ShoatID"`
}

func GetShoatByID(id int) *Short {
	s := Short{}
	result := db.First(&s, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &s
}

func InsertShoat(s Short) {
	db.Create(&s)
}

func DeleatShoat(id int) {
	s := Short{}
	db.Delete(&s, id)
}
