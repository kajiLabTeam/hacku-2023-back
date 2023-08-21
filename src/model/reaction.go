package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Reaction struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         string `gorm:"type:varchar(28)" json:"userId"`
	ShortID        int    `json:"shotId"`
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

func InsertReaction(r Reaction) error {
	result := db.Create(&r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteReaction(uid string, sId int, rId int) error {
	r := Reaction{}
	result := db.Where("user_id = ? AND short_id = ? AND reaction_list_id = ?", uid, sId, rId).Delete(&r)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("No matching data found to delete")
	}
	return nil
}
