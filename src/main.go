package main

import (
	"github.com/kajiLabTeam/hacku-2023-back/model"
	//"github.com/kajiLabTeam/hacku-2023-back/router"
)

func main() {
	//router.Init()
	model.CreateAllTabale()
	model.InsertTestData()
}
