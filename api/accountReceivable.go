package api

import (
	"jea-api/controller"
	"jea-api/models"
	"jea-api/repository"

	"github.com/gin-gonic/gin"
)

// NewAccountReceivable setup account receivable API
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
