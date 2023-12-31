package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetShort(c *gin.Context) {
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

	sId := c.Param("shortId")
	if sId != "" {
		sIdInt, err := strconv.Atoi(sId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid shortId",
			})
			return
		}
		s := model.GetShortByID(sIdInt)
		if s != nil && s.GenreID != 0 {
			sl := []Slide{}
			for i := 0; i < len(model.GetSlideByShortID(s.ID)); i++ {
				url, _ := integrations.GetFileUrl(model.GetSlideByShortID(s.ID)[i].Voice)
				tmp := Slide{
					Script:   model.GetSlideByShortID(s.ID)[i].Script,
					Content:  model.GetSlideByShortID(s.ID)[i].SlideText,
					VoiceURL: url,
				}
				sl = append(sl, tmp)
			}
			t := []string{}
			for i := 0; i < len(model.GetTagByShortID(s.ID)); i++ {
				t = append(t, model.GetTagByShortID(s.ID)[i].TagName)
			}

			r := []Reaction{}
			rl := model.GetReactionList()
			for i := 0; i < len(rl); i++ {
				u_r := model.GetReactionByShortID(s.ID, rl[i].ID)
				count := 0
				reac := false
				for j := 0; j < len(u_r); j++ {
					if u_r[j].UserID == uid {
						count++
					}
				}
				if count == 1 {
					reac = true
				}
				tmp := Reaction{
					Count:   len(model.GetReactionByShortID(s.ID, rl[i].ID)),
					Reacted: reac}
				r = append(r, tmp)
			}
			print(r[0].Count)
			rs := Reactions{
				Heart: r[0],
				Good:  r[1],
				Smile: r[2],
			}
			// 日にちまでのフォーマット
			dateFormat := "2006-01-02"
			fDate := s.CreatedAt.Format(dateFormat)
			result := Presentation{
				ID:        sIdInt,
				Title:     s.Title,
				Speaker:   s.Speaker,
				Slides:    sl,
				Tags:      t,
				Genre:     model.GetGenreByID(s.GenreID).GenreName,
				Views:     len(model.GetBrowsingHistoryByShortID(s.ID)),
				Poster:    model.GetUserByID(s.UserID).UserName,
				CreatedAt: fDate,
				Reactions: rs,
			}
			//出力
			c.JSON(http.StatusOK, gin.H{"shorts": result})
			model.InsertBrowsingHistory(model.BrowsingHistory{UserID: uid, ShortID: sIdInt, ReadAt: time.Now()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Unknown short id",
			})
			return
		}
	} else {
		s := model.GetRandomShort()
		//作った構造体にデータを入れる
		result := []Presentation{}
		for i := 0; i < len(s); i++ {
			sl := []Slide{}
			for j := 0; j < len(model.GetSlideByShortID(s[i].ID)); j++ {
				url, _ := integrations.GetFileUrl(model.GetSlideByShortID(s[i].ID)[j].Voice)
				tmp := Slide{
					Script:   model.GetSlideByShortID(s[i].ID)[j].Script,
					Content:  model.GetSlideByShortID(s[i].ID)[j].SlideText,
					VoiceURL: url,
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
			tmp := Presentation{
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
			model.InsertBrowsingHistory(model.BrowsingHistory{UserID: uid, ShortID: tmp.ID, ReadAt: time.Now()})
			result = append(result, tmp)
		}
		//出力
		c.JSON(http.StatusOK, gin.H{"shorts": result})
	}
}
