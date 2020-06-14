package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/common"
)

func health(c *gin.Context) {
	c.JSON(200, common.JSON{
		"status": "OK",
	})
}

func NewHealthAPI(router *gin.RouterGroup) {
	var api = router.Group("/health")
	{
		api.GET("", health)
	}
}
