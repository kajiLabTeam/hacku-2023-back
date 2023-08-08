package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	KeywordID []Keyword `gorm:"foreignkey:ID"`
	ShoatID   []Shoat   `gorm:"foreignkey:ID"`
}
