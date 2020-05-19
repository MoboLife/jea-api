package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jea-api/permissions"
)

var PASSWORD = "7962g97f2b02gf7826t8f2g828vy0f0682gy80f"

func GenerateToken(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func SignToken(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(PASSWORD))
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte(PASSWORD), nil
	})
}

func AuthCheckMiddleware(c *gin.Context) {
	var tokenString string
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		tokenString = c.GetHeader("Authorization")
	}
	if tokenString == "" {
		c.JSON(401, JSON{"message": "Token not found", "code": 9})
		c.Abort()
		return
	}
	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(401, JSON{"message": "Invalid token", "code": 8})
		c.Abort()
		return
	}
	permissions.PermissionMiddleware(token.Claims, c)
}
