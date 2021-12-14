package main

import "github.com/gin-gonic/gin"

func pingHandler(ginCtx *gin.Context) {
	ginCtx.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	server := gin.Default()

	server.GET("/ping", pingHandler)

	server.Run()
}
