package api

import (
	"jea-api/common"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, common.JSON{
		"status": "OK",
	})
}

// NewHealthAPI create health for API
func NewHealthAPI(router *gin.RouterGroup) {
	var api = router.Group("/health")
	{
		api.GET("", health)
	}
}
