package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

type EmployerAPI struct {

}

func (e *EmployerAPI) EmployerSales(ctx *gin.Context) {

}

func NewEmployerAPI(router *gin.RouterGroup) {
	var employer EmployerAPI
	var api = router.Group("/employers")
	var ginController = controller.NewGinController(&models.Employer{})
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Sales")},
			Find: []repository.Options{repository.WithPreloads("Sales")},
		})
		api.GET("/sales", employer.EmployerSales)
	}
}
