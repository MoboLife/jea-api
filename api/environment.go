package api

import (
	"jea-api/common"
	"jea-api/controller"
	"jea-api/database"
	"jea-api/environment"
	"jea-api/models"
	"jea-api/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EnvironmentAPI api for environment
type EnvironmentAPI struct {
	EnvironmentController controller.EnvironmentController
	EnvironmentRepository repository.Repository
}

func (e *EnvironmentAPI) setupRepository(c *gin.Context) {
	if e.EnvironmentRepository == nil {
		var db = database.GetDatabase(c)
		e.EnvironmentRepository = repository.NewRepository(&models.Environment{}, db)
		e.EnvironmentController = controller.NewEnvironmentController(db)
	}
}

func (e *EnvironmentAPI) createEnvironment(c *gin.Context) {
	var idStr = c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}
	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		c.Status(400)
		return
	}
	item, err := e.EnvironmentRepository.Find(id)
	if err != nil {
		c.Status(404)
		return
	}
	env := item.(*models.Environment)
	err = e.EnvironmentController.Create(env.EID)
	if err != nil {
		c.JSON(400, common.JSON{"message": "Error in create environment", "error": err.Error()})
		return
	}
	env.Created = true
	err = e.EnvironmentRepository.Update(&env, id)
	if err != nil {
		c.JSON(400, common.JSON{"message": "Error in update environment"})
		return
	}
	userRepository := repository.NewRepository(&models.User{}, environment.UseEnvironment(env.EID, database.GetDatabase(c)))
	err = userRepository.Create(&user)
	if err != nil {
		c.JSON(400, common.JSON{"message": "Error in create user in environment"})
		return
	}
	c.Status(201)
}

func (e *EnvironmentAPI) updateEnvironment(ctx *gin.Context) {
	var idStr = ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.Status(400)
		return
	}
	item, err := e.EnvironmentRepository.Find(id)
	if err != nil {
		ctx.Status(404)
		return
	}
	env := item.(*models.Environment)
	err = e.EnvironmentController.Update(env.EID)
	if err != nil {
		ctx.JSON(400, common.JSON{"message": "Error in create environment", "error": err.Error()})
		return
	}
	ctx.Status(200)
}

// NewEnvironment create API for environment
func NewEnvironment(group *gin.RouterGroup) {
	var environmentAPI = EnvironmentAPI{}
	var routerGroup = group.Group("/environment")
	var ginController = controller.NewGinController(&models.Environment{})
	{
		controller.NewGinControllerWrapper(routerGroup, ginController, true)
		routerGroup.Use(environmentAPI.setupRepository)
		routerGroup.POST("/:id/create", environmentAPI.createEnvironment)
		routerGroup.POST("/:id/update", environmentAPI.updateEnvironment)
	}
}
