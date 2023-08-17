package model

import "time"

func InsertTestData() {
	user := []User{
		{ID: "0000000000000000000000000001", UserName: "mizutani"},
		{ID: "0000000000000000000000000002", UserName: "shika"},
		{ID: "0000000000000000000000000003", UserName: "hoge"},
	}
	keyword := []Keyword{
		{ID: 1, KeywordName: "Golang"},
		{ID: 2, KeywordName: "Java"},
		{ID: 3, KeywordName: "C"},
		{ID: 4, KeywordName: "C--"},
		{ID: 5, KeywordName: "C++"},
		{ID: 6, KeywordName: "MySQL"},
	}
	genre := []Genre{
		{ID: 1, GenreName: "プログラミング言語"},
		{ID: 2, GenreName: "バックエンド"},
		{ID: 3, GenreName: "フロントエンド"},
	}
	reaction_list := []ReactionList{
		{ID: 1, ReactionName: "heart"},
		{ID: 2, ReactionName: "good"},
		{ID: 3, ReactionName: "smile"},
	}
	achievement := Achievement{ID: 1, UserID: user[0].ID, KeywordID: keyword[0].ID}
	short := []Short{
		{ID: 1, UserID: user[0].ID, GenreID: genre[1].ID, Title: "よくわからんGo", Speaker: "ずんだもん", CreatedAt: time.Now()},
		{ID: 2, UserID: user[0].ID, GenreID: genre[0].ID, Title: "C--とは？", Speaker: "あんこもん", CreatedAt: time.Now()},
		{ID: 3, UserID: user[1].ID, GenreID: genre[1].ID, Title: "MySQL完全に理解した", Speaker: "きなこもん", CreatedAt: time.Now()},
		{ID: 4, UserID: user[2].ID, GenreID: genre[0].ID, Title: "GoとJava", Speaker: "アデリーペンギン", CreatedAt: time.Now()},
		{ID: 5, UserID: user[1].ID, GenreID: genre[1].ID, Title: "Javaだぞ", Speaker: "鹿", CreatedAt: time.Now()},
	}
	tag := []Tag{
		{ID: 1, KeywordID: keyword[0].ID, ShortID: short[0].ID},
		{ID: 2, KeywordID: keyword[1].ID, ShortID: short[4].ID},
		{ID: 3, KeywordID: keyword[2].ID, ShortID: short[0].ID},
		{ID: 4, KeywordID: keyword[3].ID, ShortID: short[1].ID},
		{ID: 5, KeywordID: keyword[4].ID, ShortID: short[0].ID},
		{ID: 6, KeywordID: keyword[5].ID, ShortID: short[0].ID},
		{ID: 7, KeywordID: keyword[0].ID, ShortID: short[3].ID},
		{ID: 8, KeywordID: keyword[1].ID, ShortID: short[3].ID},
		{ID: 9, KeywordID: keyword[2].ID, ShortID: short[3].ID},
	}
	slide := []Slide{
		{ID: 1, ShortID: short[0].ID, SlideText: "GOがよくわからん", VoiceURL: "yyy/yyy", PageNumber: 1, Script: "ゴーがよくわからねえのだ"},
		{ID: 2, ShortID: short[0].ID, SlideText: "わからぬ", VoiceURL: "yyy/yyy", PageNumber: 2, Script: "わからぬ"},
		{ID: 3, ShortID: short[3].ID, SlideText: "GOとjavaを比べるよ", VoiceURL: "yyy24/yyyh", PageNumber: 1, Script: "クエー!!"},
		{ID: 4, ShortID: short[3].ID, SlideText: "なんか違うね", VoiceURL: "ydyy/ytyy", PageNumber: 2, Script: "クエー?"},
	}
	browsinghistory := []BrowsingHistory{
		{ID: 1, UserID: user[0].ID, ShortID: short[0].ID, ReadAt: time.Now()},
		{ID: 2, UserID: user[1].ID, ShortID: short[3].ID, ReadAt: time.Now()},
		{ID: 3, UserID: user[0].ID, ShortID: short[3].ID, ReadAt: time.Now()},
	}
	reaction := []Reaction{
		{ID: 1, UserID: user[0].ID, ShortID: short[3].ID, ReactionListID: reaction_list[0].ID},
		{ID: 2, UserID: user[1].ID, ShortID: short[3].ID, ReactionListID: reaction_list[0].ID},
		{ID: 3, UserID: user[2].ID, ShortID: short[3].ID, ReactionListID: reaction_list[1].ID},
		{ID: 4, UserID: user[0].ID, ShortID: short[3].ID, ReactionListID: reaction_list[1].ID},
		{ID: 5, UserID: user[1].ID, ShortID: short[3].ID, ReactionListID: reaction_list[2].ID},
	}
	for i := 0; i < len(user); i++ {
		InsertUser(user[i])
	}
	for i := 0; i < len(keyword); i++ {
		InsertKeyword(keyword[i])
	}
	for i := 0; i < len(genre); i++ {
		InsertGenre(genre[i])
	}
	for i := 0; i < len(reaction_list); i++ {
		InsertReactionList(reaction_list[i])
	}

	InsertAchievement(achievement)
	for i := 0; i < len(short); i++ {
		InsertShort(short[i])
	}
	for i := 0; i < len(tag); i++ {
		InsertTag(tag[i])
	}

	for i := 0; i < len(slide); i++ {
		InsertSlide(slide[i])
	}
	for i := 0; i < len(browsinghistory); i++ {
		InsertBrowsingHistory(browsinghistory[i])
	}
	for i := 0; i < len(reaction); i++ {
		InsertReaction(reaction[i])
	}
}
