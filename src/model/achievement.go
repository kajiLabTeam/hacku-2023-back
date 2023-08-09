package model

import (
	"errors"

	"gorm.io/gorm"
)

type Achievement struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	UserID    []User    `gorm:"foreignkey:ID"`
	KeywordID []Keyword `gorm:"foreignkey:ID"`
}

func GetAchievementByID(id int) *Achievement {
	achievement := Achievement{}
	result := db.First(&achievement, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &achievement
}

func InsertAchievement(achievement Achievement) {
	db.Create(&achievement)
}

func DeleatAchievement(id int) {
	achievement := Achievement{}
	db.Delete(&achievement, id)
}
