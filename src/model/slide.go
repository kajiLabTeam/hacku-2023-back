package model

import (
	"errors"

	"gorm.io/gorm"
)

type Slide struct {
	ID         int     `gorm:"primarykey;AUTO_INCREMENT"`
	ShoatID    []Shoat `gorm:"foreignkey:ID"`
	SlideText  string
	SlideURL   string
	VoiceURL   string
	PageNumber int
}

func GetSlideByID(id int) *Slide {
	slide := Slide{}
	result := db.First(&slide, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &slide
}

func InsertSlide(slide Slide) {
	db.Create(&slide)
}

func DeleatSlide(id int) {
	slide := Slide{}
	db.Delete(&slide, id)
}
