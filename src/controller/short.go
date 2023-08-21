package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/service"
)

type ShortPost struct {
	Title   string  `json:"title"`
	Speaker string  `json:"speaker"`
	Slides  []Slide `json:"Slide"`
}
type Slide struct {
	Script string `json:"script"`
	Slide  string `json:"slide"`
}

func PutShort(c *gin.Context) {
	var req ShortPost
	h := c.Request.Header.Get("Authorization")
	a := strings.TrimPrefix(h, "Bearer ")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateShort(a, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
