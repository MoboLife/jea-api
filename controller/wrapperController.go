package controller

import (
	"jea-api/common"
	"jea-api/repository"

	"github.com/gin-gonic/gin"
)

// MethodsOptions customize options of routers
type MethodsOptions struct {
	FindAll []repository.Options
	Find    []repository.Options
	Create  []repository.Options
	Delete  []repository.Options
	Update  []repository.Options
}

// NewGinControllerWrapper wrapper for ginController
func NewGinControllerWrapper(routerGroup *gin.RouterGroup, ginController GinController, secure bool, methods ...MethodsOptions) {
	if secure {
		routerGroup.Use(common.AuthCheckMiddleware)
	}
	routerGroup.Use(ginController.SetupRepository)
	var methodOptions MethodsOptions
	if len(methods) >= 1 {
		methodOptions = methods[0]
	}
	routerGroup.GET("", func(ctx *gin.Context) {
		var options = methodOptions.FindAll
		options = append(options, repository.WithFilters(ctx, repository.LimitAndPageFilter(), repository.OrderingFilter()))
		ginController.FindAll(options...)(ctx)
	})
	routerGroup.GET("/:id", ginController.Find(methodOptions.Find...))
	routerGroup.POST("", ginController.Create(methodOptions.Create...))
	routerGroup.DELETE("/:id", ginController.Delete(methodOptions.Delete...))
	routerGroup.PUT("/:id", ginController.Update(methodOptions.Update...))
}
