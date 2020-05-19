package api

import (
	"github.com/gin-gonic/gin"
	"jea-api/common"
	"jea-api/database"
	"jea-api/models"
	"jea-api/permissions"
	"jea-api/repository"
)

type ProfileAPI struct {
	UserRepository 		repository.Repository
}

func (u *ProfileAPI) setupRepository(c *gin.Context) {
	u.UserRepository = repository.NewRepository(&models.User{}, database.GetDatabase(c))
	c.Next()
}

func (u *ProfileAPI) user(c *gin.Context) {
	userPermissionsItem, exists := c.Get("permissions")
	if !exists{
		c.Status(400)
		return
	}
	userPermissions := userPermissionsItem.(*permissions.UserPermission)
	user, err := u.UserRepository.Find(userPermissions.UserId, repository.WithPreloads("Groups"))
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, user)
}

func NewProfile(group *gin.RouterGroup) {

	var api ProfileAPI
	var apiGroup = group.Group("/profile")
	{
		apiGroup.Use(common.AuthCheckMiddleware)
		apiGroup.Use(api.setupRepository)
		apiGroup.GET("", api.user)
	}

}
