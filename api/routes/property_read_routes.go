package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/api/handlers"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
)

func PropertyReadRoutes(ging *gin.Engine, dao *dao.DAO) {

	ging.GET("/ping", func(c *gin.Context) {
		handlers.PingHandler(c, dao)
	})
}
