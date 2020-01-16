package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	. "github.com/followgo/myadmin/config"
)

// UseJWT 检查 Token 信息
func UseJWT(r echoRouter) {
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(C.HTTP.TokenSigningKey),
		ContextKey: TokenContextKey,
		SuccessHandler: func(c echo.Context) {
			// 检查用户权限
		},
	}))
}
