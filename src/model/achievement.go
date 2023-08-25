package model

import (
	"errors"

	"gorm.io/gorm"
)

type Achievement struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           string `gorm:"type:varchar(28)" json:"userId"`
	AchievementName  string `json:"achievementName"`
	AchievementImage string `json:"achievementImage"`
}

func GetAchievementByID(id int) *Achievement {
	a := Achievement{}
	result := db.First(&a, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &a
}

func GetAchievementByUserID(id string) []Achievement {
	a := []Achievement{}
	result := db.Where("user_id = ?", id).Find(&a)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return a
}

func InsertAchievement(a Achievement) {
	db.Create(&a)
}

func DeleteAchievement(id int) {
	a := Achievement{}
	db.Delete(&a, id)
}

func GetAchieveByNameUserId(name, uid string) *Achievement {
	var a Achievement
	result := db.Where("achievement_name = ? AND user_id = ?", name, uid).Find(&a)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &a
}

func UpdateAchievement(new_a, a Achievement) {
	db.Model(&a).Updates(new_a)
}
