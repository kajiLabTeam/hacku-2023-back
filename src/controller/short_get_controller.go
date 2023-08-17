package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetShort(c *gin.Context) {
	s_id_str := c.Param("shortId")
	s_id, err := strconv.Atoi(s_id_str)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid shortId",
		})
		return
	}
	s := model.GetShortByID(s_id)
	if s != nil && s.GenreID != 0 {
		type Slide struct {
			Script   string `json:"script"`
			Content  string `json:"slide"`
			VoiceURL string `json:"voiceURL"`
		}

		type Presentation struct {
			ID        int            `json:"id"`
			Title     string         `json:"title"`
			Speaker   string         `json:"speaker"`
			Slides    []Slide        `json:"slides"`
			Tags      []string       `json:"tags"`
			Genre     string         `json:"genre"`
			Views     int            `json:"views"`
			Poster    string         `json:"poster"`
			CreatedAt time.Time      `json:"createdAt"`
			Reactions map[string]int `json:"reactions"`
		}
		sl := []Slide{}
		for i := 0; i < len(model.GetSlideByShortID(s.ID)); i++ {
			tmp := Slide{
				Script:   model.GetSlideByShortID(s.ID)[i].Script,
				Content:  model.GetSlideByShortID(s.ID)[i].SlideText,
				VoiceURL: model.GetSlideByShortID(s.ID)[i].VoiceURL,
			}
			sl = append(sl, tmp)
		}
		t := []string{}
		for i := 0; i < len(model.GetTagByShortID(s.ID)); i++ {
			t = append(t, model.GetKeywordByID(model.GetTagByShortID(s.ID)[i].KeywordID).KeywordName)
		}
		r := make(map[string]int)
		rl := model.GetReactionList()
		for i := 0; i < len(rl); i++ {
			r[rl[i].ReactionName] = len(model.GetReactionByShortID(s_id, rl[i].ID))
		}
		result := Presentation{
			ID:        s_id,
			Title:     s.Title,
			Speaker:   s.Speaker,
			Slides:    sl,
			Tags:      t,
			Genre:     model.GetGenreByID(s.GenreID).GenreName,
			Views:     len(model.GetBrowsingHistoryByShortID(s.ID)),
			Poster:    model.GetUserByID(s.UserID).UserName,
			CreatedAt: s.CreatedAt,
			Reactions: r,
		}
		//出力
		c.JSON(http.StatusOK, gin.H{"shorts": result})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Unknown ShortID",
		})
	}
}
