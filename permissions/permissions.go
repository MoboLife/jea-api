package permissions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jea-api/database"
	"jea-api/environment"
	"jea-api/models"
	"jea-api/repository"
)

type Permissions int

type UserPermission struct {
	UserId            int64
	PermissionValue   int64
	GroupsPermissions []int64
}

func (u *UserPermission) SetPermission(permissions Permissions) {
	u.PermissionValue = int64(SetBit(uint64(u.PermissionValue), int(permissions)))
}

func (u *UserPermission) HasPermission(permissions Permissions) bool {
	var hasOnGroup bool
	for _, groupPermission := range u.GroupsPermissions {
		hasOnGroup = GetBit(uint64(groupPermission), int(permissions))
		if hasOnGroup {
			break
		}
	}
	return GetBit(uint64(u.PermissionValue), int(permissions)) || hasOnGroup
}

var (
	UserModifyOuthers Permissions = 1
	UserFindYourself  Permissions = 2
	UserFindOuthers   Permissions = 3
	UserDelete        Permissions = 4
	UserCreate        Permissions = 5
	GroupCreate       Permissions = 7
	GroupDelete       Permissions = 8
)

func GetBit(bits uint64, bitPosition int) bool {
	return (bits & (1 << bitPosition)) != 0
}

func SetBit(bits uint64, bitPosition int) uint64 {
	bits |= 1 << bitPosition
	return bits
}

func PermissionMiddleware(claims jwt.Claims, c *gin.Context) {
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return
	}
	env := claimsMap["environment"].(string)
	var db = environment.UseEnvironment(env, database.GetDatabase(c))
	userRepository := repository.NewRepository(&models.User{}, db)
	id := int64(claimsMap["id"].(float64))
	userItem, err := userRepository.Find(id, repository.WithPreloads("Groups"))
	user := userItem.(*models.User)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var groupsValue []int64
	for _, group := range user.Groups {
		groupsValue = append(groupsValue, group.Permission)
	}
	userPermission := UserPermission{
		UserId:            user.Id,
		PermissionValue:   user.Permissions,
		GroupsPermissions: groupsValue,
	}
	c.Set("permissions", &userPermission)
	c.Set("db", db)
}

func GetPermission(c *gin.Context) *UserPermission {
	value, exists := c.Get("permissions")
	if !exists {
		return nil
	}
	perm, ok := value.(*UserPermission)
	if !ok {
		return nil
	}
	return perm
}

type UserCustomVerification func(permission *UserPermission, c *gin.Context) bool

func CustomPrivateRoute(routeFunc func(c *gin.Context), customVerification UserCustomVerification) func(c *gin.Context) {
	return func(c *gin.Context) {
		userPermission := GetPermission(c)
		if customVerification != nil {
			if customVerification(userPermission, c) {
				routeFunc(c)
				return
			}
			c.JSON(401, map[string]interface{}{"message": "You dont have permission", "code": 6})
			return
		}
	}
}

func PrivateRoute(permission Permissions, routeFunc func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		userPermission := GetPermission(c)
		if !userPermission.HasPermission(permission) {
			c.JSON(401, map[string]interface{}{"message": "You dont have permission", "code": 6})
			return
		}
		routeFunc(c)
	}
}
