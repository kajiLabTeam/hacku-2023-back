package main

import (
	"github.com/kajiLabTeam/hacku-2023-back/model"
	"github.com/kajiLabTeam/hacku-2023-back/router"
	"github.com/kajiLabTeam/hacku-2023-back/service"
)

func main() {
	model.CreateAllTable()
	service.SetKeyword()
	router.Init()
}
