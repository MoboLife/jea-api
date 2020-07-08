package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

func NewProductAPI(router *gin.RouterGroup) {
	var api = router.Group("/products")
	{
		var ginController = controller.NewGinController(&models.Product{})
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Stock", "StockTransfers", "Group")},
		})
	}
}
