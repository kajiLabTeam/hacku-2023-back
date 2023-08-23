package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID                string            `gorm:"primarykey;type:varchar(28)" json:"id"`
	UserName          string            `json:"userName"`
	Achievements      []Achievement     `gorm:"foreignkey:UserID"`
	Shoats            []Short           `gorm:"foreignkey:UserID"`
	BrowsingHistories []BrowsingHistory `gorm:"foreignkey:UserID"`
	Reactions         []Reaction        `gorm:"foreignkey:UserID"`
}

func GetUserByID(id string) *User {
	u := User{}
	result := db.First(&u, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &u
}

func InsertUser(u User) {
	db.Create(&u)
}
