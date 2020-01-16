package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/apiv1"
	"github.com/followgo/myadmin/router/mw"
)

func RegisterAPIv1(e *echo.Echo) {
	v1 := e.Group("/v1")
	registerAPI(v1, "/hello", new(apiv1.Hello), nil, nil)

	mw.UseJWT(v1)
	registerAPI(v1, "/hello1", new(apiv1.Hello), nil, nil)
}
