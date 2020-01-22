package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/apiv1"
)

func RegisterAPIv1(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// 测试
	v1.GET("/helloworld", apiv1.HelloWorld)

	// 不需要登陆
	v1.POST("/login", new(apiv1.Login).LoginByLocal)
	v1.POST("/ldap/login", new(apiv1.Login).LoginByLDAP)
	v1.GET("/files/:uuid", new(apiv1.FileAPI).Download)
	v1.GET("/images/:uuid", new(apiv1.FileAPI).Image)

	// 需要 token 才能访问
	// mw.UsePermission(v1)
	v1.POST("/logout", new(apiv1.Login).Logout)
	v1.POST("/refresh_token", new(apiv1.Login).RefreshToken)

	v1.POST("/files", new(apiv1.FileAPI).Upload) // 上传文件
	v1.GET("/files", new(apiv1.FileAPI).Select)  // 列出文件

	registerAPI(v1, "/users", new(apiv1.UserAPI), nil, nil)
}
