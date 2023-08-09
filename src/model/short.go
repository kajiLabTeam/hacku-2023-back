package model

import (
	"time"
)

type Shoat struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	UserID    []Keyword `gorm:"foreignkey:ID"`
	GenreID   []Genre   `gorm:"foreignkey:ID"`
	Title     string
	CreatedAt time.Time
}
