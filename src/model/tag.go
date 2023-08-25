package model

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	TagName string `json:"keywordId"`
	ShortID   int `json:"shortId"`
}

func GetTagByID(id int) *Tag {
	t := Tag{}
	result := db.First(&t, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &t
}

func GetTagByShortID(id int) []Tag {
	t := []Tag{}
	result := db.Where("short_id = ?", id).Find(&t).Distinct("keyword_id")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return t
}

func GetTagByKeywordID(k_id []int) []Tag {
	t := []Tag{}
	subQuery := db.Table("tags").Select("short_id").Where("keyword_id IN (?)", k_id).
		Group("short_id").Having("COUNT(DISTINCT keyword_id) = ?", len(k_id))
	db.Where("keyword_id IN (?)", k_id).Where("short_id IN (?)", subQuery).
		Find(&t)
	return t
}

func GetTagByName(k_name []string) []Tag {
	var tags []Tag
	for _, v := range k_name {
		var tag Tag
		db.Where("tagname = ?", v).First(&tag)
		tags = append(tags, tag)
	}
	return tags
}

func InsertTag(t Tag) {
	db.Create(&t)
}

func DeleteTag(id int) {
	t := Tag{}
	db.Delete(&t, id)
}
