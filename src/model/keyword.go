package model

type Keyword struct {
	ID          int `gorm:"primarykey;AUTO_INCREMENT"`
	KeywordName string
}
