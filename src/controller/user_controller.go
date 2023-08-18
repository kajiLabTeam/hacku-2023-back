package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

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
	dates := model.GetBrowsingDayByUserID(u_id)
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
