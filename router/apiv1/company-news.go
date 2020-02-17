package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// CompanyNewsAPI 公司资讯API
type CompanyNewsAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *CompanyNewsAPI) Get(c echo.Context) error {
	news := &model.CompanyNews{UUID: c.Param("uuid")}
	has, err := news.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, news)
}

// Select 列出所有选择的对象
func (api *CompanyNewsAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	news, err := new(model.CompanyNews).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.CompanyNews).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "news": news})
}

// Create 创建一个新对象
func (api *CompanyNewsAPI) Create(c echo.Context) error {
	news := &model.CompanyNews{}
	if err := c.Bind(news); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := news.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, news)
}

// Update 完全更新一个对象
func (api *CompanyNewsAPI) Update(c echo.Context) error {
	news := &model.CompanyNews{}
	if err := c.Bind(news); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	news.UUID = c.Param("uuid")

	n, err := news.Update(nil, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, news)
}

// Patch 修改一个对象的属性
func (api *CompanyNewsAPI) Patch(c echo.Context) error {
	news := &model.CompanyNews{}
	if err := c.Bind(news); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	news.UUID = c.Param("uuid")

	cols := strings.Split(c.QueryParam("cols"), ",")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := news.Update(cols, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, news)
}

// Delete 删除一个对象
func (api *CompanyNewsAPI) Delete(c echo.Context) error {
	news := &model.CompanyNews{UUID: c.Param("uuid")}
	ok, err := news.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
