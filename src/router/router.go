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

	r.GET("/api/short/search", controller.SearchShort)
	r.GET("/api/short/get", controller.GetShort)
	r.GET("/api/short/get/:shortId", controller.GetShort)
	r.GET("/api/user/profile", controller.GetProfile)
	//r.GET("/api/user/post/history/get/", controller.)
	r.GET("/api/user/browsing/history/", controller.GetBrowsingHistory)

	r.Run(":8000")
}
