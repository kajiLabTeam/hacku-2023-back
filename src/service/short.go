package service

import (
	"log"
	"os"

	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func CreateShort(uid string, req model.ShortPost) error {
	var short model.Short
	var slides []model.Slide
	var tags []model.Tag
	var style, gid int
	t := req.Tags
	s := req.Slides

	switch req.Genre {
	case "web":
		gid = 1
	case "バックエンド":
		gid = 2
	case "モバイル":
		gid = 3
	case "インフラ":
		gid = 4
	case "ゲーム":
		gid = 5
	case "その他":
		gid = 6
	}

	switch req.Speaker {
	case "四国めたん":
		style = 2
	case "ずんだもん":
		style = 3
	case "春日部つむぎ":
		style = 8
	case "雨晴はう":
		style = 10
	case "波音リツ":
		style = 9
	case "玄野武宏":
		style = 11
	case "白上虎太郎":
		style = 12
	case "青山龍星":
		style = 13
	case "冥鳴ひまり":
		style = 14
	case "九州そら":
		style = 16
	case "もち子さん":
		style = 20
	case "剣崎雌雄":
		style = 21
	case "WhiteCUL":
		style = 23
	case "後鬼":
		style = 27
	case "No.7":
		style = 29
	case "ちび式じい":
		style = 42
	case "櫻歌ミコ":
		style = 43
	case "小夜/SAYO":
		style = 46
	case "ナースロボ＿タイプＴ":
		style = 47
	case "†聖騎士 紅桜†":
		style = 51
	case "雀松朱司":
		style = 52
	case "麒ヶ島宗麟":
		style = 53
	case "春歌ナナ":
		style = 54
	case "猫使アル":
		style = 55
	case "猫使ビィ":
		style = 58
	case "中国うさぎ":
		style = 61
	}

	for i, v := range s {
		var slide model.Slide
		b, _ := getBinary(v.Script, style)
		fn, _ := makeMp3File(b, uid)
		rp, _ := integrations.Storage(fn)
		slide.SlideText = v.Slide
		slide.Voice = rp
		slide.PageNumber = i
		slide.Script = v.Script
		slides = append(slides, slide)
		os.Remove(fn)
	}

	for _, v := range t {
		var tag model.Tag
		if k := model.GetKeywordByName(v); k != nil {
			tag.TagName = k.KeywordName
			tags = append(tags, tag)
		}
	}

	short.UserID = uid
	short.Title = req.Title
	short.GenreID = gid
	short.Slides = slides
	short.Speaker = req.Speaker
	short.Tags = tags
	if err := model.InsertShort(short); err != nil {
		log.Fatal(err)
	}

	for _, v := range tags {
		var achieve model.Achievement
		key := model.GetKeywordByName(v.TagName)
		achieve.AchievementImage = key.ImageURL
		a := model.GetAchieveByNameUserId(v.TagName, uid)
		model.UpdateAchievement(achieve, *a)
	}

	return nil
}
