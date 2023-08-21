package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func PostReaction(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(auth, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	uid := t.UID
	sId := c.Param("shortId")
	var req ReactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sIdInt, err := strconv.Atoi(sId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shortId",
		})
		return
	}

	rId, err := model.GetReactionIDByName(req.Reaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction"})
		return
	}

	log.Println(uid, sIdInt, *rId, req.Reaction)

	err = model.InsertReaction(model.Reaction{UserID: uid, ShortID: sIdInt, ReactionListID: *rId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reaction": req.Reaction})
}

func DeleteReaction(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	tId := strings.TrimPrefix(auth, "Bearer ")
	t, err := integrations.VerifyIDToken(tId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	uid := t.UID
	sId := c.Param("shortId")
	var req ReactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sIdInt, err := strconv.Atoi(sId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shortId",
		})
		return
	}

	rId, err := model.GetReactionIDByName(req.Reaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction"})
		return
	}

	fmt.Println(uid, sIdInt, *rId, *rId)

	err = model.DeleteReaction(uid, sIdInt, *rId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reaction": req.Reaction})
}
