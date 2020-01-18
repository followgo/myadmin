package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct{}

func (api *Hello) Get(c echo.Context) error {
	return c.JSON( http.StatusOK, echo.Map{
		"uuid": c.Param("uuid"),
		"text": "Hello World",
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
