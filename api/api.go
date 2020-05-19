package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/modules"
)

func NewAPI(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		modules.Build(apiGroup, Modules{})
		NewProfile(apiGroup)
		NewSale(apiGroup)
		NewLogin(apiGroup)
		NewEnvironment(apiGroup)
	}
}
