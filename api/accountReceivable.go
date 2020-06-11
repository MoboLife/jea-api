package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"
)

func NewAccountReceivable(group *gin.RouterGroup) {
	var ginController = controller.NewGinController(&models.AccountReceivable{})
	var api = group.Group("/accountReceivable")
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("Client", "Company")},
			Find:    []repository.Options{repository.WithPreloads("Client", "Company")},
		})
	}
}
