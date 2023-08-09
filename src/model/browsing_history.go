package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BrowsingHistory struct {
	ID      int     `gorm:"primarykey;AUTO_INCREMENT"`
	UserID  []User  `gorm:"foreignkey:ID"`
	ShoatID []Shoat `gorm:"foreignkey:ID"`
	ReadAt  time.Time
}

func GetBrowsingHistoryByID(id int) *BrowsingHistory {
	browsinghistory := BrowsingHistory{}
	result := db.First(&browsinghistory, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &browsinghistory
}

func InsertBrowsingHistory(browsinghistory BrowsingHistory) {
	db.Create(&browsinghistory)
}

func DeleatBrowsingHistory(id int) {
	browsinghistory := BrowsingHistory{}
	db.Delete(&browsinghistory, id)
}
