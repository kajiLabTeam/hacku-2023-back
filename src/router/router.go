package router

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/controller"
)

func Init() {
	gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.GET("/api/short/search", controller.SearchShort)
	r.GET("/api/short/get", controller.GetShort)
	r.GET("/api/short/get/:shortId", controller.GetShort)
	r.GET("/api/user/profile", controller.GetProfile)
	r.GET("/api/user/post/history/get", controller.GetPostingHistory)
	r.GET("/api/user/browsing/history", controller.GetBrowsingHistory)

	r.POST("/api/user/post", controller.PostUser)
	r.POST("/api/short/post", controller.PostShort)
	r.POST("/api/short/:shortId/reaction/add", controller.PostReaction)

	r.DELETE("/api/short/:shortId/reaction/remove", controller.DeleteReaction)

	r.Run(":8084")
}
