package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/util/errors"
	"github.com/followgo/myadmin/util/mylogrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// startHTTPServer 启动HTTP服务
func startHTTPServer() {
	e := echo.New()

	e.Debug = C.HTTP.Debug
	e.DisableHTTP2 = e.Debug // 禁用http2
	e.HideBanner = !e.Debug  // 不打印banner

	// 添加全局中间件
	addEchoMiddlewares(e)

	// 注册URL路由=>API
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	// 启动服务
	go func() {
		logrus.Infoln("HTTP服务启动...")
		if err := e.Start(C.HTTP.ListenAddr); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logrus.WithError(err).Fatalln("HTTP服务意外停止")
			}
		}
	}()

	// 捕获进程信号，等待退出
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	logrus.Infof("捕获到进程型号: %v", <-sig)

	if err := e.Close(); err != nil {
		logrus.WithError(err).Fatalln("停止 HTTP 服务")
	}
	logrus.Infoln("已经停止 HTTP 服务")
}

// addEchoMiddlewares 添加 Echo 中间件函数
func addEchoMiddlewares(e *echo.Echo) {
	if !e.Debug {
		e.Use(middleware.Recover())
	}

	// JWT Middleware
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(C.HTTP.TokenSigningKey),
	// 	ContextKey: TokenContextKey,
	// }))

	// CORS 跨域资源共享
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: C.HTTP.AllowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// 日志记录
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		Output: mylogrus.NewWriterWithSizeRotate(C.HTTP.AccessFile, 100, 100, 30),
	}))
}
