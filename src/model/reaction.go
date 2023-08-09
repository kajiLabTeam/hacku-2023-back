package model

import (
	"errors"

	"gorm.io/gorm"
)

type Reaction struct {
	ID             int            `gorm:"primarykey;AUTO_INCREMENT"`
	UserID         []User         `gorm:"foreignkey:ID"`
	ShoatID        []Shoat        `gorm:"foreignkey:ID"`
	ReactionListID []ReactionList `gorm:"foreignkey:ID"`
}

func GetReactionByID(id int) *Reaction {
	reaction := Reaction{}
	result := db.First(&reaction, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &reaction
}

func InsertReaction(reaction Reaction) {
	db.Create(&reaction)
}

func DeleatReaction(id int) {
	reaction := Reaction{}
	db.Delete(&reaction, id)
}
