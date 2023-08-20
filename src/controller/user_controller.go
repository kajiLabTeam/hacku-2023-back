package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

type Short struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Slide  string `json:"slide"`
	Views  int    `json:"views"`
	Poster string `json:"poster"`
}

func GetProfile(c *gin.Context) {
	type Achievement struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	type Genre struct {
		Name       string `json:"name"`
		Color      string `json:"color"`
		DailyViews []int  `json:"dailyViews"`
	}

	type Report struct {
		Dates  []string `json:"dates"`
		Genres []Genre  `json:"genres"`
	}
	type Data struct {
		Achievements []Achievement `json:"achievements"`
		Report       Report        `json:"report"`
	}
	//本来はトークンから取得
	u_id := "0000000000000000000000000001"
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
	for i := 0; i < len(model.GetAchievementByUserID(u_id)); i++ {
		tmp := Achievement{
			Name: model.GetKeywordByID(model.GetAchievementByUserID(u_id)[i].KeywordID).KeywordName,
			Link: model.GetKeywordByID(model.GetAchievementByUserID(u_id)[i].KeywordID).ImageURL,
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
			tmp := model.GetBrowsingHistoryByUserIDAndDay(u_id, dates[j])
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
	page := c.DefaultQuery("page", "")
	//本来はトークンから取得
	u_id := "0000000000000000000000000001"

	var result []Short
	bh := model.Get100BrowsingHistoryByUserID(u_id, page)
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
	page := c.DefaultQuery("page", "")
	//本来はトークンから取得
	u_id := "0000000000000000000000000001"

	var result []Short
	ph := model.Get100ShortByUserID(u_id, page)
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
