package model

import (
	"time"
)

type Browsing_History struct {
	ID      int     `gorm:"primarykey;AUTO_INCREMENT"`
	UserID  []User  `gorm:"foreignkey:ID"`
	ShoatID []Shoat `gorm:"foreignkey:ID"`
	ReadAt  time.Time
}
