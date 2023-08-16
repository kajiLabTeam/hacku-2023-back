package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetShort(c *gin.Context) {
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
		s = model.GetShortByID(s_id)

	} else if tags == "" && title != "" {
		//Tagsの入力なしでTitleに入力ある時はショートを全件取得
		s = model.GetAllShort()
	}

	//titleに入力があれば部分一致抽出
	if title != "" {
		var tmp []model.Short
		for i := 0; i < len(s); i++ {
			if strings.Contains(strings.ToLower(s[i].Title), title) {
				tmp = append(tmp, s[i])
			}
		}
		s = tmp
	}

	//json用に抽出したShortをもとに構造体を作成
	type ShortOutPut struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Slide    string `json:"slide"`
		VoiceURL string `json:"voiceURL"`
		Views    int    `json:"views"`
		Speaker  string `json:"speaker"`
		Poster   string `json:"poster"`
	}

	//作った構造体にデータを入れる
	result := []ShortOutPut{}
	for i := 0; i < len(s); i++ {
		s_out := ShortOutPut{
			ID:       s[i].ID,
			Title:    s[i].Title,
			Slide:    model.GetThumbnailByShortID(s[i].ID).SlideText,
			VoiceURL: model.GetThumbnailByShortID(s[i].ID).VoiceURL,
			Views:    len(model.GetBrowsingHistoryByShortID(s[i].ID)),
			Speaker:  s[i].Speaker,
			Poster:   model.GetUserByID(s[i].UserID).UserName,
		}
		result = append(result, s_out)
	}

	//出力
	c.JSON(http.StatusOK, gin.H{"shorts": result})

}
