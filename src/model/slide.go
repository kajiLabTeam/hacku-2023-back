package model

import (
	"errors"

	"gorm.io/gorm"
)

type Slide struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	ShortID   int    `json:"shortId"`
	SlideText string `json:"slideText"`
	//SlideURL   string `json:"slideUrl"`
	VoiceURL   string `json:"voiceUrl"`
	Script     string `json:"script"`
	PageNumber int    `json:"pageNumber"`
}

func GetSlideByID(id int) *Slide {
	s := Slide{}
	result := db.First(&s, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &s
}

func GetThumbnailByShortID(id int) *Slide {
	s := Slide{}
	result := db.Where("short_id = ?", id).Where("page_number = ?", 1).Find(&s)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &s
}

func InsertSlide(s Slide) {
	db.Create(&s)
}

func DeleteSlide(id int) {
	s := Slide{}
	db.Delete(&s, id)
}
