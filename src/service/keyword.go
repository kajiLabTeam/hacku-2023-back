package service

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func SetKeyword() {
	file, err := os.Open("service/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range rows {
		if i == 0 {
			continue
		}
		var k model.Keyword
		k.KeywordName = v[0]
		k.ImageURL = "image/" + v[1]
		model.InsertKeyword(k)
	}

}
