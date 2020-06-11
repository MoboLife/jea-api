package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

func NewAccountPayable(group *gin.RouterGroup) {
	var ginController = controller.NewGinController(&models.AccountPayable{})
	var api = group.Group("/accountPayable")
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Client", "Company")},
			Find:    []repository.Options{repository.WithPreloads("Client", "Company")},
		})
	}
}