package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prinarjayaoka/my-go-rest-api/api/routes"
	"github.com/prinarjayaoka/my-go-rest-api/dao"
)

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

	rgV1 := server.Group("/v1")
	{
		routes.PropertyReadRoutes(rgV1, dao)
		routes.PropertyCreateRoutes(rgV1, dao)
	}

	server.Run()
}
