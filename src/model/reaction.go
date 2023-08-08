package model

type Reaction struct {
	ID              int             `gorm:"primarykey;AUTO_INCREMENT"`
	UserID          []User          `gorm:"foreignkey:ID"`
	ShoatID         []Shoat         `gorm:"foreignkey:ID"`
	Reaction_ListID []Reaction_List `gorm:"foreignkey:ID"`
}
