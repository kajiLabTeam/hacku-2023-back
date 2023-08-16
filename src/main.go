package main

import (
	"github.com/kajiLabTeam/hacku-2023-back/model"
	"github.com/kajiLabTeam/hacku-2023-back/router"
)

func main() {
	model.CreateAllTable()
	model.InsertTestData()
	router.Init()
}
