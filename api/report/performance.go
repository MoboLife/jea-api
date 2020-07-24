package report

import (
	"jea-api/common"
	"jea-api/database"
	"jea-api/models"
	"jea-api/repository"

	"github.com/gin-gonic/gin"
)

// Performance struct for represent report
type Performance struct {
	AccountsPayableTotal    float64         `json:"accountsPayableTotal" db:"accounts_payable_total"`
	AccountsReceivableTotal float64         `json:"accountsReceivableTotal" db:"accounts_receivable_total"`
	SalesTotal              float64         `json:"salesTotal" db:"sales_total"`
	SalesDiscountTotal		float64			`json:"salesDiscountTotal" db:"sales_discount_total"`
	PurchasesTotal          float64         `json:"purchasesTotal" db:"purchases_total"`
	Company                 *models.Company `json:"company,omitempty" db:"-"`
}

// PerformanceReport endpoint for performance report
func PerformanceReport(filters repository.Filters) func (c *gin.Context){
	return func(c *gin.Context) {
		var db = filters.Apply(c, database.GetDatabase(c))
		var companies []*models.Company
		err := db.Find(&companies).Error
		if err != nil {
			common.SendError(c, err, 500)
			return
		}
		var performanceReports []*Performance
		for _, company := range companies {
			var performance Performance
			db.Model(&models.Sale{}).Select("sum(total) as sales_total, sum(discount) as sales_discount_total").Where("company_id = ?", company.ID).Group("company_id").Scan(&performance)
			db.Model(&models.Purchase{}).Select("sum(total) as purchases_total").Where("company_id = ?", company.ID).Group("company_id").Scan(&performance)
			db.Model(&models.AccountPayable{}).Select("sum(amount) as accounts_payable_total").Where("company_id = ?", company.ID).Group("company_id").Scan(&performance)
			db.Model(&models.AccountReceivable{}).Select("sum(amount) as accounts_receivable_total").Where("company_id = ?", company.ID).Group("company_id").Scan(&performance)
			performance.Company = company
			performanceReports = append(performanceReports, &performance)
		}
		c.JSON(200, performanceReports)
	}
}

var PerformanceFilters = []models.ModelFilter{
	models.CreatedFilter,
}

// NewPerformanceReport create performance report
func NewPerformanceReport(router *gin.RouterGroup) {
	router.GET("/performance", PerformanceReport(
		repository.Filters{
			repository.UseFilters(PerformanceFilters),
		}),
	)
}
