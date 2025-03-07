package auth

import (
	"errors"
	"jea-api/common"
	"jea-api/permissions"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PASSWORD password for use in jwt validation
var PASSWORD = "7962g97f2b02gf7826t8f2g828vy0f0682gy80f"

// GenerateToken function for create jwt tokens
func GenerateToken(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// SignToken sign jwt tokens
func SignToken(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(PASSWORD))
}

// ValidateToken validate jwt token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte(PASSWORD), nil
	})
}

// AuthCheckMiddleware gin middleware for check user jwt
func AuthCheckMiddleware(c *gin.Context) {
	var tokenString string
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		tokenString = c.GetHeader("Authorization")
	}
	if tokenString == "" {
		c.JSON(401, common.JSON{"message": "Token not found", "code": 9})
		c.Abort()
		return
	}
	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(401, common.JSON{"message": "Invalid token", "code": 8})
		c.Abort()
		return
	}
	c.Set("token", token)
	permissions.PermissionMiddleware(token.Claims, c)
}
