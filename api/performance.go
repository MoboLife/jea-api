package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"jea-api/auth"
	"jea-api/database"
	"jea-api/models"
	"jea-api/repository"
)

type PerformanceAPI struct {
}


var PerformanceFilter = []models.ModelFilter{
	{
		Query:    "company",
		Field:    "company_id",
		Multiple: false,
		Type:     models.Integer,
	},
	{
		Query:    "createdAt",
		Field:    "created_at",
		Multiple: true,
		Type:     models.Date,
	},
}

func UseOptions(ctx *gin.Context) *gorm.DB{
	return repository.UseOptions(database.GetDatabase(ctx), repository.WithFilters(ctx, repository.UseFilters(PerformanceFilter)))
}

type Performance struct {
	AccountsPayableTotal    float64 `json:"accountsPayableTotal"`
	AccountsReceivableTotal float64 `json:"accountsReceivableTotal"`
	SalesTotal              float64 `json:"salesTotal"`
	PurchasesTotal          float64 `json:"purchasesTotal"`
}

// localhost:8080/api/performance?filters=companyId=1,2&createdAt=2020-05-18
func (p *PerformanceAPI) Performance(ctx *gin.Context) {
	var db = UseOptions(ctx)
	var accountsPayable []models.AccountPayable
	err := db.Find(&accountsPayable).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, ctx.Error(err).JSON())
		return
	}
	var accountsReceivable []models.AccountReceivable
	err = db.Find(&accountsReceivable).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, ctx.Error(err).JSON())
		return
	}
	var sales []models.Sale
	err = db.Find(&sales).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, ctx.Error(err).JSON())
		return
	}
	var purchases []models.Purchase
	err = db.Find(&purchases).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, ctx.Error(err).JSON())
		return
	}
	var performance Performance
	for _, sale := range sales {
		performance.SalesTotal += float64(sale.Total)
	}
	for _, purchase := range purchases {
		performance.PurchasesTotal += purchase.Total
	}
	for _, payable := range accountsPayable {
		performance.AccountsPayableTotal += payable.Amount
	}
	for _, receivable := range accountsReceivable {
		performance.AccountsReceivableTotal += receivable.Amount
	}

	ctx.JSON(200, performance)
}

func NewPerformanceAPI(api *gin.RouterGroup) {
	var performance PerformanceAPI
	var router = api.Group("/performance")
	{
		router.Use(auth.AuthCheckMiddleware)
		router.GET("", performance.Performance)
	}
}



