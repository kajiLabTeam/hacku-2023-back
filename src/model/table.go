package model

import (
	"github.com/kajiLabTeam/hacku-2023-back/lib"
)

func CreateAllTabale() {
	db := lib.SqlConnect()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Keyword{})
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&Reaction_List{})
	db.AutoMigrate(&Shoat{})
	db.AutoMigrate(&Slide{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Reaction{})
	db.AutoMigrate(&Browsing_History{})
	db.AutoMigrate(&Achievement{})
}
