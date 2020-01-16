package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/apiv1"
	"github.com/followgo/myadmin/router/mw"
)

func RegisterAPIv1(e *echo.Echo) {
	v1 := e.Group("/v1")

	// 不需要登陆
	v1.POST("/login", new(apiv1.Login).LoginByLocal)

	// 需要 token 才能访问
	mw.UseJWT(v1)
	v1.POST("/login", new(apiv1.Login).Logout)
	v1.POST("/login", new(apiv1.Login).RefreshToken)
	registerAPI(v1, "/hello", new(apiv1.Hello), nil, nil)
}
