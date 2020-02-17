package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// PartnerAPI 合作伙伴API
type PartnerAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *PartnerAPI) Get(c echo.Context) error {
	p := &model.Partner{UUID: c.Param("uuid")}
	has, err := p.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, p)
}

// Select 列出所有选择的对象
func (api *PartnerAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	partners, err := new(model.Partner).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.Partner).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "partners": partners})
}

// Create 创建一个新对象
func (api *PartnerAPI) Create(c echo.Context) error {
	p := &model.Partner{}
	if err := c.Bind(p); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := p.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, p)
}

// Update 完全更新一个对象
func (api *PartnerAPI) Update(c echo.Context) error {
	p := &model.Partner{}
	if err := c.Bind(p); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	p.UUID = c.Param("uuid")

	n, err := p.Update(nil, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, p)
}

// Patch 修改一个对象的属性
func (api *PartnerAPI) Patch(c echo.Context) error {
	p := &model.Partner{}
	if err := c.Bind(p); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	p.UUID = c.Param("uuid")

	cols := strings.Split(c.QueryParam("cols"), ",")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := p.Update(cols, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, p)
}

// Delete 删除一个对象
func (api *PartnerAPI) Delete(c echo.Context) error {
	p := &model.Partner{UUID: c.Param("uuid")}
	ok, err := p.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
