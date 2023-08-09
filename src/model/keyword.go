package model

import (
	"errors"

	"gorm.io/gorm"
)

type Keyword struct {
	ID          int `gorm:"primarykey;AUTO_INCREMENT"`
	KeywordName string
}

func GetKeywordByID(id int) *Keyword {
	keyword := Keyword{}
	result := db.First(&keyword, id)
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
