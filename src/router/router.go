package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/controller"
)

func Init() {
	r := gin.Default()
	r.GET("/hoge", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/api/short/search", controller.GetShort)

	r.Run(":3000")
}