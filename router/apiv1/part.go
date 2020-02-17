package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
)

// PartAPI 内容部件API
type PartAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *PartAPI) Get(c echo.Context) error {
	part := &model.Part{UUID: c.Param("uuid")}
	has, err := part.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, part)
}

// Select 列出所有选择的对象
func (api *PartAPI) Select(c echo.Context) error { return echo.ErrNotFound }

// Create 创建一个新对象
func (api *PartAPI) Create(c echo.Context) error {
	part := &model.Part{}
	if err := c.Bind(part); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := part.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, part)
}

// Update 完全更新一个对象
func (api *PartAPI) Update(c echo.Context) error {
	part := &model.Part{}
	if err := c.Bind(part); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	part.Name = c.Param("uuid")

	n, err := part.Update()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, part)
}

// Patch 修改一个对象的属性
func (api *PartAPI) Patch(c echo.Context) error { return echo.ErrNotFound }

// Delete 删除一个对象
func (api *PartAPI) Delete(c echo.Context) error {
	part := &model.Part{Name: c.Param("uuid")}
	ok, err := part.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
