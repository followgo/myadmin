package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct{}

func (api *Hello) Get(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(echo.Map{
		"UUID": c.Param("uuid"),
		"TEXT": "Hello World",
	})
}

func (api *Hello) Select(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func (api *Hello) Create(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func (api *Hello) Update(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func (api *Hello) Patch(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

// Delete 删除
func (api *Hello) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
