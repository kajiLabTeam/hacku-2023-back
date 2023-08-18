package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BrowsingHistory struct {
	ID      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  string    `gorm:"type:varchar(28)" json:"userId"`
	ShortID int       `json:"shortId"`
	ReadAt  time.Time `json:"readAt"`
}

func GetBrowsingHistoryByUserID(id string) []BrowsingHistory {
	bh := []BrowsingHistory{}
	result := db.Where("user_id = ?", id).Find(&bh, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return bh
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

func GetBrowsingHistoryByUserIDAndDay(id string, day string) []BrowsingHistory {
	bh := []BrowsingHistory{}
	day = strings.TrimSuffix(day, " +0900")
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, day)
	if err != nil {
		fmt.Println("Error:", err)
	}
	result := db.Where("user_id =?", id).Where("DATE(read_at) = ?", t.Format("2006-01-02")).Find(&bh)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return bh
}

func GetBrowsingDayByUserID(id string) []string {
	var u_dates []time.Time
	var result []string
	db.Where("user_id =?", id).Model(&BrowsingHistory{}).Select("DISTINCT DATE(read_at) as read_date").Order("read_date desc").Limit(7).Scan(&u_dates)
	for i := len(u_dates) - 1; i >= 0; i-- {
		result = append(result, u_dates[i].Format("2006-01-02 15:04:05 +0900"))
	}
	return result
}

func InsertBrowsingHistory(bh BrowsingHistory) {
	db.Create(&bh)
}

func DeleteBrowsingHistory(id int) {
	bh := BrowsingHistory{}
	db.Delete(&bh, id)
}
