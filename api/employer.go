package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

type EmployerAPI struct {

}


func NewEmployerAPI(router *gin.RouterGroup) {
	var api = router.Group("/employers")
	var ginController = controller.NewGinController(&models.Employer{})
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Sales")},
			Find: []repository.Options{repository.WithPreloads("Sales")},
		})
	}
}
