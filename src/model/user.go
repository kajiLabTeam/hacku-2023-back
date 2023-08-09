package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primarykey;type:varchar(28)"`
	UserName string
}

func GetUserByID(id string) interface{} {
	user := User{}
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return user
}

func InsertUser(user User) {
	db.Create(&user)
}
