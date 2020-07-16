package report

import (
	"github.com/gin-gonic/gin"
	"jea-api/database"
	"jea-api/models"
)

type Performance struct {

}

func PerformanceReport(c *gin.Context) {
	var db = database.GetDatabase(c)
	db.Model(&models.Company{}).Select("")
}

func NewPerformanceReport(router *gin.RouterGroup) {
	router.GET("/performance", PerformanceReport)
}