package mw

import (
	"github.com/labstack/echo/v4/middleware"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/util/mylogrus"
)

// UseAccessLogger 访问记录
func UseAccessLogger(r echoRouter) {
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		Output: mylogrus.NewWriterWithSizeRotate(C.HTTP.AccessFile, 100, 100, 30),
	}))
}
