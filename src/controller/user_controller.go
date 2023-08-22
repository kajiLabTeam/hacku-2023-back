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

func GetProfile(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(auth, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	uid := t.UID

	var a = []Achievement{}
	colors := []string{
		"245, 101, 101, 1",
		"66, 153, 225, 1",
		"246, 173, 85, 1",
		"72, 187, 120, 1",
		"159, 122, 234, 1",
		"160, 174, 192, 1",
	}
	var g = []Genre{}
	for i := 0; i < len(model.GetAchievementByUserID(uid)); i++ {
		tmp := Achievement{
			Name: model.GetKeywordByID(model.GetAchievementByUserID(uid)[i].KeywordID).KeywordName,
			Link: model.GetKeywordByID(model.GetAchievementByUserID(uid)[i].KeywordID).ImageURL,
		}
		a = append(a, tmp)
	}
	var dates []string
	for i := 6; i >= 0; i-- {
		tmp := time.Now().AddDate(0, 0, -i).Format("2006-01-02 15:04:05 +0900")
		dates = append(dates, tmp)
	}
	all_g := model.GetAllGenre()
	for i := 0; i < len(dates); i++ {

	}
	for i := 0; i < len(all_g); i++ {
		var d_v []int
		for j := 0; j < len(dates); j++ {
			tmp := model.GetBrowsingHistoryByUserIDAndDay(uid, dates[j])
			count := 0
			for _, bh := range tmp {
				if model.GetShortByID(bh.ShortID).GenreID == all_g[i].ID {
					count++
				}
			}
			d_v = append(d_v, count)
		}

		tmp := Genre{
			Name:       all_g[i].GenreName,
			Color:      colors[i],
			DailyViews: d_v,
		}
		g = append(g, tmp)
	}
	r := Report{Dates: dates, Genres: g}
	result := Data{Achievements: a, Report: r}
	//出力
	c.JSON(http.StatusOK, result)
}

func GetBrowsingHistory(c *gin.Context) {
	h := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(h, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	uid := t.UID
	page := c.DefaultQuery("page", "")

	var result []Short
	bh := model.Get100BrowsingHistoryByUserID(uid, page)
	for i := 0; i < len(bh); i++ {
		tmp := Short{
			ID:     bh[i].ShortID,
			Title:  model.GetShortByID(bh[i].ShortID).Title,
			Slide:  model.GetThumbnailByShortID(bh[i].ShortID).SlideText,
			Views:  len(model.GetBrowsingHistoryByShortID(bh[i].ShortID)),
			Poster: model.GetUserByID(bh[i].UserID).UserName,
		}
		result = append(result, tmp)
	}
	if result == nil {
		result = []Short{}
	}
	//出力
	c.JSON(http.StatusOK, gin.H{"browsingHistories": result})

}

func GetPostingHistory(c *gin.Context) {
	h := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(h, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	uid := t.UID
	page := c.DefaultQuery("page", "")
	var result []Short
	offset, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	}
	ph := model.Get100ShortByUserID(uid, offset)
	for i := 0; i < len(ph); i++ {
		tmp := Short{
			ID:     ph[i].ID,
			Title:  ph[i].Title,
			Slide:  model.GetThumbnailByShortID(ph[i].ID).SlideText,
			Views:  len(model.GetBrowsingHistoryByShortID(ph[i].ID)),
			Poster: model.GetUserByID(ph[i].UserID).UserName,
		}
		result = append(result, tmp)
	}
	if result == nil {
		result = []Short{}
	}
	//出力
	c.JSON(http.StatusOK, gin.H{"postingHistories": result})
}

func PostUser(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(auth, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	uid := t.UID
	u := model.GetUserByID(uid)
	if u != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User already exists"})
	}

	model.InsertUser(model.User{ID: uid, UserName: ""})

	c.JSON(http.StatusOK, gin.H{"id": uid, "name": ""})
}
