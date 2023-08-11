package model

import (
	"github.com/kajiLabTeam/hacku-2023-back/lib"
)

var db = lib.SqlConnect()

func CreateAllTabale() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Keyword{})
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&ReactionList{})
	db.AutoMigrate(&Short{})
	db.AutoMigrate(&Slide{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Reaction{})
	db.AutoMigrate(&BrowsingHistory{})
	db.AutoMigrate(&Achievement{})
}
