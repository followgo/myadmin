package mw

import (
	"github.com/labstack/echo/v4"
)

// echoRouter echo.Echo, echo.Group
type echoRouter interface {
	Use(middleware ...echo.MiddlewareFunc)
}
