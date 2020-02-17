package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// SolutionAPI 解决方案API
type SolutionAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *SolutionAPI) Get(c echo.Context) error {
	solution := &model.Solution{UUID: c.Param("uuid")}
	has, err := solution.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, solution)
}

// Select 列出所有选择的对象
func (api *SolutionAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	solutions, err := new(model.Solution).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.Solution).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "solutions": solutions})
}

// Create 创建一个新对象
func (api *SolutionAPI) Create(c echo.Context) error {
	solution := &model.Solution{}
	if err := c.Bind(solution); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := solution.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, solution)
}

// Update 完全更新一个对象
func (api *SolutionAPI) Update(c echo.Context) error {
	solution := &model.Solution{}
	if err := c.Bind(solution); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	solution.UUID = c.Param("uuid")

	n, err := solution.Update(nil, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, solution)
}

// Patch 修改一个对象的属性
func (api *SolutionAPI) Patch(c echo.Context) error {
	solution := &model.Solution{}
	if err := c.Bind(solution); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	solution.UUID = c.Param("uuid")

	cols := strings.Split(c.QueryParam("cols"), ",")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := solution.Update(cols, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, solution)
}

// Delete 删除一个对象
func (api *SolutionAPI) Delete(c echo.Context) error {
	solution := &model.Solution{UUID: c.Param("uuid")}
	ok, err := solution.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
