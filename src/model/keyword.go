package model

import (
	"errors"

	"gorm.io/gorm"
)

type Keyword struct {
	ID          int         `gorm:"primaryKey;autoIncrement" json:"id"`
	KeywordName string      `json:"keywordName"`
	ImageURL    string      `json:"imageUrl"`
	Achievement Achievement `gorm:"foreignkey:KeywordID"`
	Tags        []Tag       `gorm:"foreignkey:KeywordID"`
}

func GetKeywordByID(id int) *Keyword {
	k := Keyword{}
	result := db.Find(&k, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &k
}

func GetKeywordByName(name string) *Keyword {
	k := Keyword{}
	result := db.Where("keyword_name = ?", name).First(&k)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &k
}

func InsertKeyword(k Keyword) {
	db.Create(&k)
}

func DeleteKeyword(id int) {
	keyword := Keyword{}
	db.Delete(&keyword, id)
}
