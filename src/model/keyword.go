package model

import (
	"errors"

	"gorm.io/gorm"
)

type Keyword struct {
	ID          int         `gorm:"primaryKey;autoIncrement" json:"id"`
	KeywordName string      `json:"keywordName"`
	Achievement Achievement `gorm:"foreignkey:KeywordID"`
	Tags        []Tag       `gorm:"foreignkey:KeywordID"`
}

func GetKeywordByID(id int) *Keyword {
	k := Keyword{}
	result := db.First(&k, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &k
}

func GetKeywordByName(name string) *Keyword {
	k := Keyword{}
	result := db.First(&k, name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &k
}

func InsertKeyword(k Keyword) {
	db.Create(&k)
}

func DeleatKeyword(id int) {
	keyword := Keyword{}
	db.Delete(&keyword, id)
}
