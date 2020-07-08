package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

func NewCarRentalAPI(router *gin.RouterGroup) {
	var api = router.Group("/carRentals")
	{
		var ginController = controller.NewGinController(&models.CarRental{})
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Client")},
			Find:    []repository.Options{repository.WithPreloads("Client")},
		})
	}
}
