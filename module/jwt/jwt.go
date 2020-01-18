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

	// Set claims
	mapClaims := jwt.MapClaims(claims)
	now := time.Now()
	mapClaims["exp"] = now.Add(30 * time.Minute).Unix()
	mapClaims["nbf"] = now.Add(-1 * time.Second).Unix()
	mapClaims["iat"] = now.Unix()

	// Create token
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	// Generate encoded token
	tokenStr, err := myToken.SignedString([]byte(TokenSigningKey))
	return TokenAuthScheme + " " + tokenStr, err
}

// GetClaimsFromToken 从上下文中提取 Claims
func GetClaimsFromToken(c echo.Context) map[string]interface{} {
	myToken := c.Get(TokenContextKey).(*jwt.Token)
	return myToken.Claims.(jwt.MapClaims)
}
