package model

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	KeywordID []Keyword `gorm:"foreignkey:ID"`
	ShoatID   []Shoat   `gorm:"foreignkey:ID"`
}

func GetTagByID(id int) *Tag {
	tag := Tag{}
	result := db.First(&tag, id)
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
