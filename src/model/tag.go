package model

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	KeywordID int `json:"keywordId"`
	ShoatID   int `json:"shoatId"`
}

func GetTagByID(id int) *Tag {
	t := Tag{}
	result := db.First(&t, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &t
}

func GetTagByKeywordID(keyword_id int) *Tag {
	t := Tag{}
	result := db.Where("keywordId = ?", keyword_id).First(&t)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &t
}

func InsertTag(t Tag) {
	db.Create(&t)
}

func DeleteTag(id int) {
	t := Tag{}
	db.Delete(&t, id)
}
