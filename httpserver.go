package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/router"
	"github.com/followgo/myadmin/util/errors"
)

// startHTTPServer 启动HTTP服务
func startHTTPServer() {
	e := echo.New()
	e.DisableHTTP2 = true // 禁用http2，前端一般使用 nginx

	e.Debug = C.HTTP.Debug
	e.HideBanner = !e.Debug // 不打印banner

	// 添加全局中间件
	router.AddGlobalMiddlewares(e)

	// 注册URL路由=>API
	router.RegisterAPIv1(e)

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
