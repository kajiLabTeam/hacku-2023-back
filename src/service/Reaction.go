package service

import "github.com/kajiLabTeam/hacku-2023-back/model"

func SetReactionList() {
	rl := []model.ReactionList{{ReactionName: "heart"}, {ReactionName: "good"}, {ReactionName: "smile"}}
	for _, v := range rl {
		model.InsertReactionList(v)
	}
}
