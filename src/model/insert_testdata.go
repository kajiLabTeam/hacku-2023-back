package model

import "time"

func InsertTestData() {
	user := User{ID: "0000000000000000000000000001", UserName: "mizutani"}
	keyword := Keyword{ID: 1, KeywordName: "Golang"}
	genre := Genre{ID: 1, GenreName: "プログラミング言語"}
	reaction_list := ReactionList{ID: 1, ReactionName: "good"}
	achievement := Achievement{ID: 1, UserID: user.ID, KeywordID: keyword.ID}
	short := Short{ID: 1, UserID: user.ID, GenreID: genre.ID, Title: "よくわからんGo", CreatedAt: time.Now()}
	tag := Tag{ID: 1, KeywordID: keyword.ID, ShoatID: short.ID}
	slide := Slide{ID: 1, ShoatID: short.ID, SlideText: "hoge", SlideURL: "xxx/xxx", VoiceURL: "yyy/yyy", PageNumber: 1, Script: "台本"}
	browsinghistory := BrowsingHistory{ID: 1, UserID: user.ID, ShoatID: short.ID, ReadAt: time.Now()}
	reaction := Reaction{ID: 1, UserID: user.ID, ShoatID: short.ID, ReactionListID: reaction_list.ID}
	InsertUser(user)
	InsertKeyword(keyword)
	InsertGenre(genre)
	InsertReactionList(reaction_list)
	InsertAchievement(achievement)
	InsertShoat(short)
	InsertTag(tag)
	InsertSlide(slide)
	InsertBrowsingHistory(browsinghistory)
	InsertReaction(reaction)
}
