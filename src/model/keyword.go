package model

import (
	"errors"

	"gorm.io/gorm"
)

type Keyword struct {
	ID          int         `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	KeywordName string      `json:"keywordName"`
	Achievement Achievement `gorm:"foreignkey:KeywordID"`
	Tags        []Tag       `gorm:"foreignkey:KeywordID"`
}

func GetKeywordByID(id int) *Keyword {
	keyword := Keyword{}
	result := db.First(&keyword, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &keyword
}

func GetKeywordByName(name string) *Keyword {
	keyword := Keyword{}
	result := db.First(&keyword, name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &keyword
}

func InsertKeyword(keyword Keyword) {
	db.Create(&keyword)
}

func DeleatKeyword(id int) {
	keyword := Keyword{}
	db.Delete(&keyword, id)
}
