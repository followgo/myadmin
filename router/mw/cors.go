package mw

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	. "github.com/followgo/myadmin/config"
)

// UseCORS 允许跨域访问
func UseCORS(r echoRouter) {
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: C.HTTP.AllowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		MaxAge:       3600,
	}))
}
