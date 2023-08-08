package model

type Slide struct {
	ID         int     `gorm:"primarykey;AUTO_INCREMENT"`
	ShoatID    []Shoat `gorm:"foreignkey:ID"`
	SlideText  string
	SlideURL   string
	VoiceURL   string
	PageNumber int
}
