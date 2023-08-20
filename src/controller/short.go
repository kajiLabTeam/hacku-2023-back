package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
	"github.com/kajiLabTeam/hacku-2023-back/service"
)

func PostShort(c *gin.Context) {
	var req model.Short
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
