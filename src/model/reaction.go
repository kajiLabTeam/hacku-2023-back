package model

import (
	"errors"

	"gorm.io/gorm"
)

type Reaction struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         string `gorm:"type:varchar(28)" json:"userId"`
	ShortID        int    `json:"shoatId"`
	ReactionListID int    `json:"reactionListId"`
}

func GetReactionByID(id int) *Reaction {
	r := Reaction{}
	result := db.First(&r, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &r
}

func GetReactionByShortID(s_id int, rl_id int) []Reaction {
	r := []Reaction{}
	result := db.Where("short_id = ?", s_id).Where("reaction_list_id = ?", rl_id).Find(&r)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return r
}

func InsertReaction(r Reaction) {
	db.Create(&r)
}

func DeleteReaction(id int) {
	r := Reaction{}
	db.Delete(&r, id)
}
