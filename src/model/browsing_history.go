package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BrowsingHistory struct {
	ID      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  string    `gorm:"type:varchar(28)" json:"userId"`
	ShoatID int       `json:"shoatId"`
	ReadAt  time.Time `json:"readAt"`
}

func GetBrowsingHistoryByID(id int) *BrowsingHistory {
	bh := BrowsingHistory{}
	result := db.First(&bh, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &bh
}

func InsertBrowsingHistory(bh BrowsingHistory) {
	db.Create(&bh)
}

func DeleatBrowsingHistory(id int) {
	bh := BrowsingHistory{}
	db.Delete(&bh, id)
}
