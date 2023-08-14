package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func GetShort(c *gin.Context) model.Short {
	tags := c.Query("tags")
	t := strings.Split(tags, ",")
	print(t[0])
	print(t[1])
	k := model.GetKeywordByName(t[0])
	if k != nil {
		t_id := model.GetTagByKeywordID(k.ID)
		s := model.GetShoatByID(t_id.ID)
		return s
	} else {
		var s *model.Short
		s = nil
		return s
	}
}
