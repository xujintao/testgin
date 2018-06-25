package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/testgin/models"
)

func Like(ctx *gin.Context) {
	var like models.Like
	ctx.BindJSON(&like)
	if err := models.DBWriteLike(&like); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"uid":    "1001",
		"result": "success",
	})
}

func LikeInfo(ctx *gin.Context) {
	strUid := ctx.Query("uid")
	uid, _ := strconv.ParseUint(strUid, 10, 64)
	likes := models.DBReadLikeByUid(uint(uid))
	ctx.JSON(http.StatusOK, likes)
}
