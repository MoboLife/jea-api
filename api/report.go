package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"jea-api/auth"
	"jea-api/common"
	"jea-api/database"
	"jea-api/models"
)

type ReportAPI struct {

}

type BestSeller struct {
	SellerId		int64				`json:"sellerId" db:"seller_id"`
	SalesTotal		float64				`json:"salesTotal" db:"sales_total"`
	SalesCount		int64				`json:"salesCount" db:"sales_count"`
	Seller			*models.Employer	`json:"seller,omitempty"`
}

func (r *ReportAPI) BestSellers(c *gin.Context) {
	var db = database.GetDatabase(c)
	var bestSellers []*BestSeller
	db.LogMode(true)
	err := db.Model(&models.Sale{}).Select("seller_id, sum(total) as sales_total, count(seller_id) as sales_count").Where("seller_id IS NOT NULL").Group("seller_id").Order("sales_total desc").Limit(10).Scan(&bestSellers).Error
	if err != nil {
		common.SendError(c, err, 500)
		return
	}
	for _, seller := range bestSellers {
		seller.Seller = &models.Employer{}
		err = db.First(seller.Seller, seller.SellerId).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			common.SendError(c, err, 500)
			return
		}
	}
	c.JSON(200, bestSellers)
}


func NewReportAPI(router *gin.RouterGroup) {
	var report = ReportAPI{}
	var api = router.Group("/reports")
	{
		api.Use(auth.AuthCheckMiddleware)
		api.GET("/bestsellers", report.BestSellers)
	}
}