package mw

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/module/jwt"
	"github.com/followgo/myadmin/module/onlineuser"
)

// UsePermission 权限许可管理
func UsePermission(r echoRouter) {
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(TokenSigningKey),
		ContextKey:    TokenContextKey,
		TokenLookup:   TokenLookup,
		AuthScheme:    TokenAuthScheme,
	}))

	r.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := jwt.GetClaimsFromToken(c)
			user := onlineuser.GetUser(claims["uuid"].(string))

			if user.Username == "" {
				return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "unknown username", Internal: nil}
			}

			return next(c)
		}
	})
}
