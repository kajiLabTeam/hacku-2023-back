package model

type Achievement struct {
	ID        int       `gorm:"primarykey;AUTO_INCREMENT"`
	UserID    []User    `gorm:"foreignkey:ID"`
	KeywordID []Keyword `gorm:"foreignkey:ID"`
}
