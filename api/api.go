package api

import (
	"jea-api/modules"

	"github.com/gin-gonic/gin"
)

// NewAPI setup application API
func NewAPI(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		NewHealthAPI(apiGroup)
		modules.Build(apiGroup, Modules{})
		NewProfile(apiGroup)
		NewSale(apiGroup)
		NewLogin(apiGroup)
		NewEnvironment(apiGroup)
		NewAccountPayable(apiGroup)
		NewAccountReceivable(apiGroup)
		NewSessionAPI(apiGroup)
		NewPurchaseAPI(apiGroup)
		NewPerformanceAPI(apiGroup)
		NewProductAPI(apiGroup)
		NewCarRentalAPI(apiGroup)
	}
}
