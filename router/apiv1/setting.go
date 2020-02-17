package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
)

// SettingAPI 网站设置API
type SettingAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *SettingAPI) Get(c echo.Context) error {
	s := &model.Setting{UUID: c.Param("uuid")}
	has, err := s.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, s)
}

// Select 列出所有选择的对象
func (api *SettingAPI) Select(c echo.Context) error { return echo.ErrNotFound }

// Create 创建一个新对象
func (api *SettingAPI) Create(c echo.Context) error {
	s := &model.Setting{}
	if err := c.Bind(s); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := s.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, s)
}

// Update 完全更新一个对象
func (api *SettingAPI) Update(c echo.Context) error {
	s := &model.Setting{}
	if err := c.Bind(s); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	s.Name = c.Param("uuid")

	n, err := s.Update()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, s)
}

// Patch 修改一个对象的属性
func (api *SettingAPI) Patch(c echo.Context) error { return echo.ErrNotFound }

// Delete 删除一个对象
func (api *SettingAPI) Delete(c echo.Context) error {
	s := &model.Setting{Name: c.Param("uuid")}
	ok, err := s.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
