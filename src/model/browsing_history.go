package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BrowsingHistory struct {
	ID      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  string    `gorm:"type:varchar(28)" json:"userId"`
	ShortID int       `json:"shortId"`
	ReadAt  time.Time `json:"readAt"`
}

func GetBrowsingHistoryByID(id int) []BrowsingHistory {
	bh := []BrowsingHistory{}
	result := db.Find(&bh, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return bh
}

func GetBrowsingHistoryByShortID(id int) []BrowsingHistory {
	bh := []BrowsingHistory{}
	result := db.Where("short_id = ?", id).Find(&bh)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return bh
}

func InsertBrowsingHistory(bh BrowsingHistory) {
	db.Create(&bh)
}

func DeleteBrowsingHistory(id int) {
	bh := BrowsingHistory{}
	db.Delete(&bh, id)
}
