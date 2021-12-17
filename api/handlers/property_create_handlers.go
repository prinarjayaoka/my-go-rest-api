package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
	"github.com/prinarjayaoka/my-go-rest-api/models"
)

func CreateDataHandler(ginCtx *gin.Context, dao *dao.DAO) {
	var t models.Thread

	errJson := ginCtx.ShouldBindJSON(&t)
	if errJson != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "ERROR",
			"data":    errJson.Error(),
		})
		return
	}

	affectedRows, err := dao.ThreadStore.CreateThread(&models.Thread{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
	})

	if affectedRows > 0 {
		ginCtx.JSON(201, gin.H{
			"status":  "SUCCESS",
			"message": "SUCCESS",
			"data":    nil,
		})
	} else {
		ginCtx.JSON(409, gin.H{
			"status":  "ERROR",
			"message": "ERROR",
			"data":    err.Error(),
		})
	}
}
