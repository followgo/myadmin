package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/followgo/myadmin/config"
)

// GenerateTokenString 生成 token 字符串
func GenerateTokenString(claims map[string]interface{}, lifetime time.Duration) (string, error) {
	// Create token
	myToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	mapClaims := myToken.Claims.(jwt.MapClaims)
	for k, v := range claims {
		mapClaims[k] = v
	}
	claims["exp"] = time.Now().Add(lifetime).Unix()

	// Generate encoded token
	return myToken.SignedString([]byte(config.TokenSigningKey))
}

// GetClaims 从上下文中提取 Claims
func GetClaims(c echo.Context) (map[string]interface{}, error) {
	myToken, ok := c.Get(config.TokenContextKey).(*jwt.Token)
	if !ok {
		return nil, errors.New("illegal token string")
	}

	claims, ok := myToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims error")
	}

	return claims, nil
}
