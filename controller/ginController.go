package controller

import (
	"jea-api/database"
	"jea-api/repository"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GinController controller for models and gin
type GinController interface {
	FindAll(options ...repository.Options) func(ctx *gin.Context)
	Find(options ...repository.Options) func(ctx *gin.Context)
	Create(options ...repository.Options) func(ctx *gin.Context)
	Delete(options ...repository.Options) func(ctx *gin.Context)
	Update(options ...repository.Options) func(ctx *gin.Context)
	Patch(options ...repository.Options) func(ctx *gin.Context)
	SetupRepository(ctx *gin.Context)
}

// GinControllerContext context of gin controller
type GinControllerContext struct {
	Model      interface{}
	ModelType  reflect.Type
	Repository repository.Repository
}

// SetupRepository middleware for setup repository
func (g *GinControllerContext) SetupRepository(ctx *gin.Context) {
	g.Repository = repository.NewRepository(g.Model, database.GetDatabase(ctx))
}

// FindAll gin router for findAll items of model
func (g *GinControllerContext) FindAll(options ...repository.Options) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		options = append(options, repository.WithFilters(ctx, repository.LimitAndPageFilter()))
		items, err := g.Repository.FindAll(options...)
		if err != nil {
			ctx.Status(400)
			return
		}
		total, err := g.Repository.Total()
		if err != nil {
			ctx.Status(400)
			return
		}
		ctx.JSON(200, FindAllResponse{
			Items: items,
			Total: total,
		})
	}
}

// Find gin router for find item of model
func (g *GinControllerContext) Find(options ...repository.Options) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var idStr = ctx.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.Status(400)
			return
		}
		item, err := g.Repository.Find(id, options...)
		if err != nil {
			if err.Error() == "record not found" {
				ctx.Status(404)
				return
			}
			ctx.Status(400)
			return
		}
		if item == nil {
			ctx.Status(404)
			return
		}
		ctx.JSON(200, item)
	}
}

// Create gin router for create item of model
func (g *GinControllerContext) Create(options ...repository.Options) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var entity = reflect.New(g.ModelType).Interface()
		err := ctx.BindJSON(entity)
		if err != nil {
			ctx.Status(400)
			return
		}
		err = g.Repository.Create(entity, options...)
		if err != nil {
			ctx.Status(400)
			return
		}
		ctx.JSON(201, entity)
	}
}

// Delete gin router for delete item of model
func (g *GinControllerContext) Delete(options ...repository.Options) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var idStr = ctx.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.Status(400)
			return
		}
		err = g.Repository.Delete(id, options...)
		if err != nil {
			ctx.Status(400)
			return
		}
		ctx.Status(200)
	}

}

// Update gin router for update item of model
func (g *GinControllerContext) Update(options ...repository.Options) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var idStr = ctx.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.Status(400)
			return
		}
		var entity = reflect.New(g.ModelType).Interface()
		err = ctx.BindJSON(entity)
		if err != nil {
			ctx.Status(400)
			return
		}
		err = g.Repository.Update(entity, id, options...)
		if err != nil {
			ctx.Status(400)
			return
		}
		ctx.JSON(200, entity)
	}
}

// Patch gin router for patch update (not implemented)
func (g *GinControllerContext) Patch(options ...repository.Options) func(ctx *gin.Context) {
	panic("implement me")
}

// NewGinController create gin controller
func NewGinController(model interface{}) GinController {
	var modelType = reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	return &GinControllerContext{ModelType: modelType, Model: model}
}

// FindAllResponse response for more one items
type FindAllResponse struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}
