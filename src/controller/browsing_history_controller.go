package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetBrowsingHistory(c *gin.Context) {
	page := c.DefaultQuery("page", "")
	//本来はトークンから取得
	u_id := "0000000000000000000000000001"
	type BrowsingShort struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Slide  string `json:"slide"`
		Views  int    `json:"views"`
		Poster string `json:"poster"`
	}
	var result []BrowsingShort
	bh := model.Get100BrowsingHistoryByUserID(u_id, page)
	for i := 0; i < len(bh); i++ {
		tmp := BrowsingShort{
			ID:     bh[i].ShortID,
			Title:  model.GetShortByID(bh[i].ShortID).Title,
			Slide:  model.GetThumbnailByShortID(bh[i].ShortID).SlideText,
			Views:  len(model.GetBrowsingHistoryByShortID(bh[i].ShortID)),
			Poster: model.GetUserByID(bh[i].UserID).UserName,
		}
		result = append(result, tmp)
	}
	if result == nil {
		result = []BrowsingShort{}
	}
	//出力
	c.JSON(http.StatusOK, gin.H{"browsingHistories": result})

}
