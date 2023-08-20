package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/model"
	"github.com/kajiLabTeam/hacku-2023-back/service"
)

func PutShort(c *gin.Context) {
	var req model.Short
	a := c.Request.Header.Get("Authorization")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateShort(a, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// service.slide
	c.JSON(http.StatusOK, req)
}
