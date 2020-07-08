package api

import (
	"jea-api/auth"
	"jea-api/database"
	"jea-api/models"
	"jea-api/permissions"
	"jea-api/repository"

	"github.com/gin-gonic/gin"
)

// ProfileAPI api for profiles
type ProfileAPI struct {
	UserRepository repository.Repository
}

func (u *ProfileAPI) setupRepository(c *gin.Context) {
	u.UserRepository = repository.NewRepository(&models.User{}, database.GetDatabase(c))
	c.Next()
}

func (u *ProfileAPI) user(c *gin.Context) {
	userPermissionsItem, exists := c.Get("permissions")
	if !exists {
		c.Status(400)
		return
	}
	userPermissions := userPermissionsItem.(*permissions.UserPermission)
	user, err := u.UserRepository.Find(userPermissions.UserID, repository.WithPreloads("Groups"))
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, user)
}

// NewProfile create profile API
func NewProfile(group *gin.RouterGroup) {

	var api ProfileAPI
	var apiGroup = group.Group("/profile")
	{
		apiGroup.Use(auth.AuthCheckMiddleware)
		apiGroup.Use(api.setupRepository)
		apiGroup.GET("", api.user)
	}

}
