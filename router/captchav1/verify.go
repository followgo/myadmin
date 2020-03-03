package captchav1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyHandler 检验验证码
func VerifyHandler(c echo.Context) error {
	config := new(configJsonBody)
	if err := c.Bind(config); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	if store.Verify(config.Id, config.VerifyValue, true) {
		return c.NoContent(http.StatusOK)
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "验证码不正确"}
}
