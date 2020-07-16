package report

import (
	"github.com/gin-gonic/gin"
	"jea-api/auth"
)

func NewReport(router *gin.RouterGroup) {
	var reports = router.Group("/reports")
	{
		reports.Use(auth.AuthCheckMiddleware)
		NewPerformanceReport(reports)
		NewBestSellerReport(reports)
	}
}
