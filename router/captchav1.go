package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/captchav1"
)

// RegisterCaptchaAPIv1  注册 API 到指定的 URL
func RegisterCaptchaAPIv1(e *echo.Echo) {
	e.GET("/captcha/v1", captchav1.GenerateHandler)
	e.POST("/captcha/v1", captchav1.VerifyHandler)
}
