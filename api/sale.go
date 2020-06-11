package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/database"
	"jea-api/models"
	"jea-api/repository"
	"strconv"
)

type SaleAPI struct {
	SaleRepository repository.Repository
}

func (s *SaleAPI) setupRepository(ctx *gin.Context) {
	if s.SaleRepository == nil {
		s.SaleRepository = repository.NewRepository(&models.Sale{}, database.GetDatabase(ctx))
	}
}

func (s *SaleAPI) findProducts(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.Status(400)
		return
	}
	saleItems, err := s.SaleRepository.Find(id, repository.WithPreloads("Purchaser", "Products", "Company"), repository.WithFilters(ctx, repository.LimitAndPageFilter()))
	sale := saleItems.(*models.Sale)
	if err != nil {
		ctx.Status(404)
		return
	}
	ctx.JSON(200, sale.Products)
}

func NewSale(group *gin.RouterGroup) {
	var saleAPI = SaleAPI{}
	var routerGroup = group.Group("/sales")
	var ginController = controller.NewGinController(&models.Sale{})
	{
		controller.NewGinControllerWrapper(routerGroup, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Purchaser", "Company", "Products", "Seller", "Products.Product") },
			Find:    []repository.Options{repository.WithPreloads("Purchaser", "Company", "Products", "Seller", "Products.Product")},
		})
		routerGroup.Use(saleAPI.setupRepository)
		routerGroup.GET("/:id/products", saleAPI.findProducts)
	}
}

