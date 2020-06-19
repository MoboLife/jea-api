package api

import (
	"jea-api/common"
	"jea-api/database"
	"jea-api/environment"
	"jea-api/models"
	"jea-api/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload payload for login
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	EID      string `json:"eid"`
}

// LoginAPI for api login
type LoginAPI struct {
	UserRepository repository.Repository
}

func (l *LoginAPI) setupRepository(c *gin.Context) {
	c.Next()
}

func (l *LoginAPI) login(c *gin.Context) {
	var loginPayload LoginPayload
	err := c.BindJSON(&loginPayload)
	if err != nil {
		c.Status(400)
		return
	}
	var db = environment.UseEnvironment(loginPayload.EID, database.GetDatabase(c))
	var userRepository = repository.NewRepository(&models.User{}, db)
	userItem, err := userRepository.FindAll(repository.WithWhere("username = ?", loginPayload.Username))
	if err != nil {
		c.JSON(200, common.JSON{"message": "Password wrong or user not exists", "code": -1})
		return
	}
	usersAddr := userItem.(*[]models.User)
	users := *usersAddr
	if len(users) == 0 {
		c.JSON(200, common.JSON{"message": "Password wrong or user not exists", "code": -1})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Hash), []byte(loginPayload.Password))
	if err != nil {
		c.JSON(200, common.JSON{"message": "Password wrong or user not exists", "code": -1})
		return
	}
	token := common.GenerateToken(jwt.MapClaims{
		"id":          int64(users[0].ID),
		"createdAt":   time.Now(),
		"environment": loginPayload.EID,
	})
	tokenString, err := common.SignToken(token)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie("access_token", tokenString, 360, "/", "localhost", false, false)
	c.JSON(200, common.JSON{"message": "Successful login", "code": "2", "token": tokenString})
}

// NewLogin create login API
func NewLogin(group *gin.RouterGroup) {

	var login = LoginAPI{}
	{
		group.Use(login.setupRepository)
		group.POST("/login", login.login)
	}
}
