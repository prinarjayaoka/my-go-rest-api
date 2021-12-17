package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/api/handlers"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
)

func PropertyCreateRoutes(ging *gin.Engine, dao *dao.DAO) {

	ging.POST("/create", func(c *gin.Context) {
		handlers.CreateDataHandler(c, dao)
	})
}
