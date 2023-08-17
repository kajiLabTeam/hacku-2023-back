package model

import (
	"errors"

	"gorm.io/gorm"
)

type ReactionList struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	ReactionName string     `json:"reactionName"`
	Reactions    []Reaction `gorm:"foreignkey:ReactionListID"`
}

func GetReactionList() []ReactionList {
	rl := []ReactionList{}
	result := db.Find(&rl)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return rl
}

func GetReactionListByID(id int) *ReactionList {
	rl := ReactionList{}
	result := db.First(&rl, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &rl
}

func InsertReactionList(rl ReactionList) {
	db.Create(&rl)
}

func DeleteReactionList(id int) {
	rl := ReactionList{}
	db.Delete(&rl, id)
}
