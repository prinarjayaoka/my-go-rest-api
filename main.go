package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
	"github.com/prinarjayaoka/my-go-rest-api/models"
)

func pingHandler(ginCtx *gin.Context, dao *dao.DAO) {
	threads, _ := dao.ThreadStore.Threads()

	ginCtx.JSON(200, gin.H{
		"status":  "SUCCESS",
		"message": "SUCCESS",
		"data":    threads,
	})
}

func createDataHandler(ginCtx *gin.Context, dao *dao.DAO) {
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

func main() {
	dsnString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true",
		"benjo",
		"12qwaszx@321123,1+1=2,h@h4H1H!",
		"localhost:3306",
		"redditgo",
	)

	dao, _ := dao.NewStore(dsnString)
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		pingHandler(c, dao)
	})

	server.POST("/create", func(c *gin.Context) {
		createDataHandler(c, dao)
	})

	server.Run()
}
