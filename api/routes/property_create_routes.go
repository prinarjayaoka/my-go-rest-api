package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/api/handlers"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
)

func PropertyCreateRoutes(rg *gin.RouterGroup, dao *dao.DAO) {

	rg.POST("/create", func(c *gin.Context) {
		handlers.CreateDataHandler(c, dao)
	})
}
