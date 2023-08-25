package service

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetAllKeyword() ([]string, []string) {
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
	var keyword, image []string
	for i, v := range rows {
		if i == 0 {
			continue
		}
		keyword = append(keyword, v[0])
		image = append(image, v[1])
	}
	return keyword, image
}

func SetKeyword() {
	keyword, image := GetAllKeyword()
	for i, v := range keyword {
		var k model.Keyword
		k.KeywordName = v
		k.ImageURL = "image/" + image[i]
		model.InsertKeyword(k)
	}
}
