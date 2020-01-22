package apiv1

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
	"github.com/followgo/myadmin/util"
)

// UserAPI 管理员用户
type UserAPI struct {
	UUID          string    `json:"uuid" form:"uuid" query:"uuid"`
	Email         string    `json:"email" form:"email" query:"email" validate:"email"`
	Username      string    `json:"username" form:"username" query:"username" validate:"required"`
	Password      string    `json:"password" form:"password" query:"password" validate:"required"`
	Roles         []string  `json:"roles" form:"roles" query:"roles"`
	Enabled       bool      `json:"enabled" form:"enabled" query:"enabled"`
	LastLoginFrom string    `json:"last_login_from" form:"last_login_from" query:"last_login_from"`
	LastLoginAt   time.Time `json:"last_login_at" form:"last_login_at" query:"last_login_at"`
	LoginCount    uint      `json:"login_count" form:"login_count" query:"login_count"`
	Created       time.Time `json:"created" form:"created" query:"created"`
	Updated       time.Time `json:"updated" form:"updated" query:"updated"`
}

// Get 根据 `ID` 获取指定的对象
func (api *UserAPI) Get(c echo.Context) error {
	user := model.User{UUID: c.Param("uuid")}
	has, err := user.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	_ = util.StructToStruct(api, user)
	return c.JSON(http.StatusOK, api)
}

// Select 列出所有选择的对象
func (api *UserAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	users, err := new(model.User).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	list := make([]UserAPI, len(users))
	for i := range users {
		_ = util.StructToStruct(&list[i], users[i])
	}

	// 数量
	total, err := new(model.User).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "data": list})
}

// Create 创建一个新对象
func (api *UserAPI) Create(c echo.Context) error {
	if err := c.Bind(api); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	user := model.User{}
	_ = util.StructToStruct(&user, api)

	ok, err := user.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	_ = util.StructToStruct(api, user)
	return c.JSON(http.StatusOK, api)
}

// Update 完全更新一个对象
func (api *UserAPI) Update(c echo.Context) error {
	if err := c.Bind(api); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	api.UUID = c.Param("uuid")

	user := model.User{}
	_ = util.StructToStruct(&user, api)

	n, err := user.Update(nil, []string{"last_login_from", "last_login_at", "login_count"})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	_ = util.StructToStruct(api, user)
	return c.JSON(http.StatusOK, api)
}

// Patch 修改一个对象的属性
func (api *UserAPI) Patch(c echo.Context) error {
	if err := c.Bind(api); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	api.UUID = c.Param("uuid")

	user := model.User{}
	_ = util.StructToStruct(&user, api)

	cols := strings.Split(c.QueryParam("cols"), ";")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := user.Update(cols, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	_ = util.StructToStruct(api, user)
	return c.JSON(http.StatusOK, api)
}

// Delete 删除一个对象
func (api *UserAPI) Delete(c echo.Context) error {
	user := &model.User{UUID: c.Param("uuid")}
	ok, err := user.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
