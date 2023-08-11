package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID               string            `gorm:"primarykey;type:varchar(28)" json:"id"`
	UserName         string            `json:"userName"`
	Achievements     []Achievement     `gorm:"foreignkey:UserID"`
	Shoats           []Short           `gorm:"foreignkey:UserID"`
	BrowsingHistorys []BrowsingHistory `gorm:"foreignkey:UserID"`
	Reactions        []Reaction        `gorm:"foreignkey:UserID"`
}

func GetUserByID(id string) *User {
	user := User{}
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func InsertUser(user User) {
	db.Create(&user)
}
