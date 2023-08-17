package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Short struct {
	ID                int               `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            string            `gorm:"type:varchar(28)" json:"userId"`
	GenreID           int               `json:"genreId"`
	Title             string            `json:"title"`
	Speaker           string            `json:"speaker"`
	CreatedAt         time.Time         `json:"createdAt"`
	BrowsingHistories []BrowsingHistory `gorm:"foreignkey:ShortID"`
	Slides            []Slide           `gorm:"foreignkey:ShortID"`
	Tags              []Tag             `gorm:"foreignkey:ShortID"`
}

func GetShortByID(id int) *Short {
	s := Short{}
	result := db.Find(&s, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &s
}

func GetShortByIDArray(id []int) []Short {
	s := []Short{}
	result := db.Where("id IN (?)", id).Find(&s).Distinct("id")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return s
}
func GetAllShort() []Short {
	s := []Short{}
	result := db.Find(&s)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return s
}
func InsertShort(s Short) {
	db.Create(&s)
}

func DeleatShoat(id int) {
	s := Short{}
	db.Delete(&s, id)
}
