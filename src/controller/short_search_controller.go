package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func SearchShort(c *gin.Context) {
	var uid string

	auth := c.Request.Header.Get("Authorization")
	if auth != "" {
		tId := strings.TrimPrefix(auth, "Bearer ")
		t, err := integrations.VerifyIDToken(tId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		uid = t.UID
	} else {
		uid = "Not logged user"
	}

	tags := c.DefaultQuery("tags", "")
	title := c.DefaultQuery("title", "")
	var s []model.Short
	if tags != "" {
		//Tagsを","で区切る
		t := strings.Split(tags, ",")

		//入力されたTagから存在するTagだけ抜き出し
		var k_name []string
		for i := 0; i < len(t); i++ {
			k := model.GetKeywordByName(t[i])
			if k != nil {
				k_name = append(k_name, k.KeywordName)
			}
		}

		//ないと思うけどTagが重複したときは１つにする
		m := make(map[string]bool)
		u_k := []string{}

		for _, ele := range k_name {
			if !m[ele] {
				m[ele] = true
				u_k = append(u_k, ele)
			}
		}

		//KeywordIDからTagを抽出
		tag := model.GetTagByName(k_name)

		//TagIDからShortを抽出
		var sId []int
		for i := 0; i < len(tag); i++ {
			sId = append(sId, tag[i].ShortID)
		}
		s = model.GetShortByIDArray(sId)

	} else if tags == "" && title != "" {
		//Tagsの入力なしでTitleに入力ある時
		s = model.GetShortByTitle(title)
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

	//作った構造体にデータを入れる
	result := []ShortOutPut{}
	for i := 0; i < len(s); i++ {
		sl := []Slide{}
		for j := 0; j < len(model.GetSlideByShortID(s[i].ID)); j++ {
			tmp := Slide{
				Script:   model.GetSlideByShortID(s[i].ID)[j].Script,
				Content:  model.GetSlideByShortID(s[i].ID)[j].SlideText,
				VoiceURL: model.GetSlideByShortID(s[i].ID)[j].Voice,
			}
			sl = append(sl, tmp)
		}
		t := []string{}
		for j := 0; j < len(model.GetTagByShortID(s[i].ID)); j++ {
			t = append(t, model.GetTagByShortID(s[i].ID)[j].TagName)
		}
		r := []Reaction{}
		rl := model.GetReactionList()
		for j := 0; j < len(rl); j++ {
			u_r := model.GetReactionByShortID(s[i].ID, rl[j].ID)
			count := 0
			reac := false
			for l := 0; l < len(u_r); l++ {
				if u_r[l].UserID == uid {
					count++
				}
			}
			if count == 1 {
				reac = true
			}
			tmp := Reaction{
				Count:   len(model.GetReactionByShortID(s[i].ID, rl[j].ID)),
				Reacted: reac}
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
