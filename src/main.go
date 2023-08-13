package main

import (
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func main() {
	model.CreateAllTable()
	model.InsertKeyword(model.Keyword{KeywordName: "shika"})
	//print(model.GetKeywordByID(2))
}
