package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	. "github.com/followgo/myadmin/config"
)

// GenerateTokenString 生成 token 字符串
func GenerateTokenString(claims map[string]interface{}, lifetime time.Duration) (string, error) {
	if claims == nil {
		claims = make(map[string]interface{})
	}

	// Create token
	myToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	mapClaims := myToken.Claims.(jwt.MapClaims)
	for k, v := range claims {
		mapClaims[k] = v
	}
	claims["exp"] = time.Now().Add(lifetime).Unix()

	// Generate encoded token
	return myToken.SignedString([]byte( C.HTTP.TokenSigningKey))
}

// GetClaimsFromToken 从上下文中提取 Claims
func GetClaimsFromToken(c echo.Context) map[string]interface{} {
	myToken := c.Get(TokenContextKey).(*jwt.Token)
	return myToken.Claims.(jwt.MapClaims)
}
