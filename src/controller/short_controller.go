package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
	"github.com/kajiLabTeam/hacku-2023-back/service"
)

func PostShort(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(auth, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	uid := t.UID
	var req model.ShortPost

	

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateShort(uid, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
