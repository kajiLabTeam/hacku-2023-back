package model

type Reaction_List struct {
	ID           int `gorm:"primarykey;AUTO_INCREMENT"`
	ReactionName string
}
