package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func SearchShort(c *gin.Context) {
	tags := c.DefaultQuery("tags", "")
	title := c.DefaultQuery("title", "")
	var s []model.Short
	if tags != "" {
		//Tagsを","で区切る
		t := strings.Split(tags, ",")

		//入力されたTagから存在するTagだけ抜き出し
		var k_id []int
		for i := 0; i < len(t); i++ {
			k := model.GetKeywordByName(t[i])
			if k != nil {
				k_id = append(k_id, k.ID)
			}
		}

		//ないと思うけどTagが重複したときは１つにする
		m := make(map[int]bool)
		u_k := []int{}

		for _, ele := range k_id {
			if !m[ele] {
				m[ele] = true
				u_k = append(u_k, ele)
			}
		}

		//KeywordIDからTagを抽出
		tag := model.GetTagByKeywordID(u_k)

		//TagIDからShortを抽出
		var s_id []int
		for i := 0; i < len(tag); i++ {
			s_id = append(s_id, tag[i].ShortID)
		}
		s = model.GetShortByIDArray(s_id)

	} else if tags == "" && title != "" {
		//Tagsの入力なしでTitleに入力ある時はショートを全件取得
		s = model.GetAllShort()
	}

	//titleに入力があれば部分一致抽出
	if title != "" {
		var tmp []model.Short
		for i := 0; i < len(s); i++ {
			if strings.Contains(strings.ToLower(s[i].Title), strings.ToLower(title)) {
				tmp = append(tmp, s[i])
			}
		}
		s = tmp
	}

	//json用に抽出したShortをもとに構造体を作成

	type Slide struct {
		Script  string `json:"script"`
		Content string `json:"slide"`
		Voice   string `json:"voiceURL"`
	}

	type Reaction struct {
		Count   int  `json:"count"`
		Reacted bool `json:"reacted"`
	}
	type Reactions struct {
		Heart Reaction `json:"heart"`
		Good  Reaction `json:"good"`
		Smile Reaction `json:"smile"`
	}

	type ShortOutPut struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Speaker   string    `json:"speaker"`
		Slides    []Slide   `json:"slides"`
		Tags      []string  `json:"tags"`
		Genre     string    `json:"genre"`
		Views     int       `json:"views"`
		Poster    string    `json:"poster"`
		CreatedAt string    `json:"createdAt"`
		Reactions Reactions `json:"reactions"`
	}

	//作った構造体にデータを入れる
	result := []ShortOutPut{}
	for i := 0; i < len(s); i++ {
		sl := []Slide{}
		for j := 0; j < len(model.GetSlideByShortID(s[i].ID)); j++ {
			tmp := Slide{
				Script:  model.GetSlideByShortID(s[i].ID)[j].Script,
				Content: model.GetSlideByShortID(s[i].ID)[j].SlideText,
				Voice:   model.GetSlideByShortID(s[i].ID)[j].Voice,
			}
			sl = append(sl, tmp)
		}
		t := []string{}
		for j := 0; j < len(model.GetTagByShortID(s[i].ID)); j++ {
			t = append(t, model.GetKeywordByID(model.GetTagByShortID(s[i].ID)[j].KeywordID).KeywordName)
		}
		r := []Reaction{}
		rl := model.GetReactionList()
		for j := 0; j < len(rl); j++ {
			tmp := Reaction{
				Count:   len(model.GetReactionByShortID(s[i].ID, rl[j].ID)),
				Reacted: true}
			r = append(r, tmp)
		}
		rs := Reactions{
			Heart: r[0],
			Good:  r[1],
			Smile: r[2],
		}
		// 日にちまでのフォーマット
		dateFormat := "2006-01-02"
		fDate := s[i].CreatedAt.Format(dateFormat)
		tmp := ShortOutPut{
			ID:        s[i].ID,
			Title:     s[i].Title,
			Speaker:   s[i].Speaker,
			Slides:    sl,
			Tags:      t,
			Genre:     model.GetGenreByID(s[i].GenreID).GenreName,
			Views:     len(model.GetBrowsingHistoryByShortID(s[i].ID)),
			Poster:    model.GetUserByID(s[i].UserID).UserName,
			CreatedAt: fDate,
			Reactions: rs,
		}
		result = append(result, tmp)
	}
	//出力
	c.JSON(http.StatusOK, gin.H{"shorts": result})

}
