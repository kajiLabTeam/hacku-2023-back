package model

import (
	"errors"

	"gorm.io/gorm"
)

type Achievement struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string `gorm:"type:varchar(28)" json:"userId"`
	KeywordID int    `json:"keywordId"`
}

func GetAchievementByID(id int) *Achievement {
	a := Achievement{}
	result := db.First(&a, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &a
}

func InsertAchievement(a Achievement) {
	db.Create(&a)
}

func DeleteAchievement(id int) {
	a := Achievement{}
	db.Delete(&a, id)
}
