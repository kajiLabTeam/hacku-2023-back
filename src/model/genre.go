package model

type Genre struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT"`
	GenreName string
}
