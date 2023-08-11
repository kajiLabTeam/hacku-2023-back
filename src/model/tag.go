package model

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	KeywordID int `json:"keywordId"`
	ShoatID   int `json:"shoatId"`
}

func GetTagByID(id int) *Tag {
	tag := Tag{}
	result := db.First(&tag, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &tag
}

func GetTagByKeywordID(keyword_id int) *Tag {
	tag := Tag{}
	result := db.Where("keywordId = ?", keyword_id).First(&tag)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &tag
}

func InsertTag(tag Tag) {
	db.Create(&tag)
}

func DeleatTag(id int) {
	tag := Tag{}
	db.Delete(&tag, id)
}
