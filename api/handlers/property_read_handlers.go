package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
)

func PingHandler(ginCtx *gin.Context, dao *dao.DAO) {
	threads, _ := dao.ThreadStore.Threads()

	ginCtx.JSON(200, gin.H{
		"status":  "SUCCESS",
		"message": "SUCCESS",
		"data":    threads,
	})
}
