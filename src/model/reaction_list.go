package model

import (
	"errors"

	"gorm.io/gorm"
)

type ReactionList struct {
	ID           int        `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	ReactionName string     `json:"reactionName"`
	Reactions    []Reaction `gorm:"foreignkey:ReactionListID"`
}

func GetReactionListByID(id int) *ReactionList {
	reactionlist := ReactionList{}
	result := db.First(&reactionlist, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &reactionlist
}

func InsertReactionList(reactionlist ReactionList) {
	db.Create(&reactionlist)
}

func DeleatReactionList(id int) {
	reactionlist := ReactionList{}
	db.Delete(&reactionlist, id)
}
