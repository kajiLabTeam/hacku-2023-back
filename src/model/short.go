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

type ShortPost struct {
	Title   string  `json:"title"`
	Speaker string  `json:"speaker"`
	Slides  []SlidePost `json:"slides"`
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

func GetShortByTitle(title string) []Short {
	s := []Short{}
	result := db.Where("title LIKE ?", "%"+title+"%").Find(&s)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return s
}

func Get100ShortByUserID(id string, offset int) []Short {
	s := []Short{}

	offset = (offset - 1) * 100
	result := db.Where("user_id =?", id).Limit(100).Offset(offset).Find(&s)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return s
}

func GetRandomShort() []Short {
	s := []Short{}
	result := db.Order("RAND()").Limit(10).Find(&s)
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
