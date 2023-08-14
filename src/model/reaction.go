package model

import (
	"errors"

	"gorm.io/gorm"
)

type Reaction struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         string `gorm:"type:varchar(28)" json:"userId"`
	ShoatID        int    `json:"shoatId"`
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

func InsertReaction(r Reaction) {
	db.Create(&r)
}

func DeleteReaction(id int) {
	r := Reaction{}
	db.Delete(&r, id)
}
