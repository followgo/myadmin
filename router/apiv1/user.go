package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// UserAPI 管理员用户API
type UserAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *UserAPI) Get(c echo.Context) error {
	user := &model.Admin{UUID: c.Param("uuid")}
	has, err := user.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, user)
}

// Select 列出所有选择的对象
func (api *UserAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	users, err := new(model.Admin).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.Admin).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "users": users})
}

// Create 创建一个新对象
func (api *UserAPI) Create(c echo.Context) error {
	user := &model.Admin{}
	if err := c.Bind(user); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := user.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, user)
}

// Update 完全更新一个对象
func (api *UserAPI) Update(c echo.Context) error {
	user := &model.Admin{}
	if err := c.Bind(user); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	user.UUID = c.Param("uuid")

	n, err := user.Update(nil, []string{"last_login_from", "last_login_at", "login_count"})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, user)
}

// Patch 修改一个对象的属性
func (api *UserAPI) Patch(c echo.Context) error {
	user := &model.Admin{}
	if err := c.Bind(user); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	user.UUID = c.Param("uuid")

	cols := strings.Split(c.QueryParam("cols"), ",")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := user.Update(cols, []string{"last_login_from", "last_login_at", "login_count"})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, user)
}

// Delete 删除一个对象
func (api *UserAPI) Delete(c echo.Context) error {
	user := &model.Admin{UUID: c.Param("uuid")}
	ok, err := user.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
