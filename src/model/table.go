package model

import (
	"github.com/kajiLabTeam/hacku-2023-back/lib"
)

var db = lib.SqlConnect()

func CreateAllTable() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Keyword{})
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&ReactionList{})
	db.AutoMigrate(&Short{})
	db.AutoMigrate(&Slide{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Reaction{})
	db.AutoMigrate(&BrowsingHistory{})

	gname := []string{
		"web",
		"バックエンド",
		"モバイル",
		"インフラ",
		"ゲーム",
		"その他",
	}
	for _, v := range gname {
		var genre Genre
		genre.GenreName = v
		InsertGenre(genre)
	}

}
