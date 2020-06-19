package api

import (
	"jea-api/common"
	"jea-api/controller"
	"jea-api/database"
	"jea-api/models"
	"jea-api/permissions"
	"jea-api/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SessionAPI api for provide session data
type SessionAPI struct {
}

// SessionRegister type for regsiter device
type SessionRegister struct {
	DeviceID string `json:"deviceId"`
	Model    string `json:"model"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

// MobileSession create session for mobile devices
func (s *SessionAPI) MobileSession(ctx *gin.Context) {
	perm, exists := ctx.Get("permissions")
	if !exists {
		ctx.AbortWithStatus(400)
		return
	}
	tk, _ := ctx.Get("token")
	token := tk.(*jwt.Token)
	permission := perm.(*permissions.UserPermission)
	var register SessionRegister
	err := ctx.BindJSON(&register)
	if err != nil {
		ctx.AbortWithStatusJSON(400, common.JSON{
			"code":    -1,
			"message": "Invalid body",
		})
	}
	var db = database.GetDatabase(ctx)
	var session models.Session
	err = db.Model(&models.Session{}).Preload("Access").Preload("User").First(&session, "token = ?", token.Raw).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.AbortWithStatus(400)
		return
	}
	if err == gorm.ErrRecordNotFound {
		session = models.Session{
			UserID: permission.UserID,
			Access: []*models.SessionAccess{
				{
					IPAddress: ctx.ClientIP(),
					AccessAt:  time.Now(),
				},
			},
			Token:    token.Raw,
			Model:    register.Model,
			DeviceID: register.DeviceID,
			Platform: register.Platform,
			Version:  register.Version,
			Type:     models.MobileSession,
		}
		err = db.Create(&session).Error
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		ctx.JSON(200, session)
		return
	}
	var access = &models.SessionAccess{
		SessionID: session.ID,
		IPAddress: ctx.ClientIP(),
		AccessAt:  time.Now(),
	}
	err = db.Create(&access).Error
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	session.Access = append(session.Access, access)
	ctx.JSON(200, session)
}

// NewSessionAPI create api for sessions
func NewSessionAPI(router *gin.RouterGroup) {
	var sessionAPI = SessionAPI{}
	api := router.Group("/session")
	var ginController = controller.NewGinController(&models.Session{})
	{
		controller.NewGinControllerWrapper(api, ginController, true, controller.MethodsOptions{
			FindAll: []repository.Options{repository.WithPreloads("User", "Access")},
			Find:    []repository.Options{repository.WithPreloads("User", "Access")},
		})
		api.POST("/mobile", sessionAPI.MobileSession)
	}
}
