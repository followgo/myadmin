package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/followgo/myadmin/router/mw"
)

// AddGlobalMiddlewares 添加 Echo 全局中间件
func AddGlobalMiddlewares(e *echo.Echo) {
	// 日志记录
	mw.UseAccessLogger(e)

	// 静态文件
	// e.Use(middleware.Static("/static"))

	// CORS 跨域资源共享
	mw.UseCORS(e)

	if !e.Debug {
		e.Use(middleware.Recover()) // 从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给 HTTPErrorHandler 处理。
		e.Use(middleware.Secure())  // 阻止跨站脚本攻击(XSS)，内容嗅探，点击劫持，不安全链接等其他代码注入攻击。
	}
}
