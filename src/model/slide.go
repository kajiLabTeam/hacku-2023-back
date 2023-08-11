package model

import (
	"errors"

	"gorm.io/gorm"
)

type Slide struct {
	ID         int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	ShoatID    int    `json:"shoatId"`
	SlideText  string `json:"slideText"`
	SlideURL   string `json:"slideUrl"`
	VoiceURL   string `json:"voiceUrl"`
	Script     string `json:"script"`
	PageNumber int    `json:"pageNumber"`
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
