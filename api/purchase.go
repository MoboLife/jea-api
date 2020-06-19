package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)


func NewPurchaseAPI(group *gin.RouterGroup) {
	var api = group.Group("/purchases")
	var ginController = controller.NewGinController(&models.Purchase{})
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Company", "Seller", "Products", "Products.Product")},
			Find:    []repository.Options{repository.WithPreloads("Company", "Seller", "Products", "Products.Product")},
		})
	}
}
