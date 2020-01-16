package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	. "github.com/followgo/myadmin/config"
)

// GenerateTokenString 生成 token 字符串
func GenerateTokenString(claims map[string]interface{}) (string, error) {
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
	now := time.Now()
	claims["exp"] = now.Add(10 * time.Minute).Unix()
	claims["nbf"] = now.Add(-1 * time.Second).Unix()
	claims["iat"] = now.Unix()

	// Generate encoded token
	return myToken.SignedString([]byte( C.HTTP.TokenSigningKey))
}

// GetClaimsFromToken 从上下文中提取 Claims
func GetClaimsFromToken(c echo.Context) map[string]interface{} {
	myToken := c.Get(TokenContextKey).(*jwt.Token)
	return myToken.Claims.(jwt.MapClaims)
}
