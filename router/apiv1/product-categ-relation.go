package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// ProductCategRelationAPI 产品和类别的关系API
type ProductCategRelationAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *ProductCategRelationAPI) Get(c echo.Context) error {
	rel := &model.ProductCategRelation{UUID: c.Param("uuid")}
	has, err := rel.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, rel)
}

// Select 列出所有选择的对象
func (api *ProductCategRelationAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	relations, err := new(model.ProductCategRelation).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.ProductCategRelation).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "relations": relations})
}

// Create 创建一个新对象
func (api *ProductCategRelationAPI) Create(c echo.Context) error {
	rel := &model.ProductCategRelation{}
	if err := c.Bind(rel); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := rel.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, rel)
}

// Update 完全更新一个对象
func (api *ProductCategRelationAPI) Update(c echo.Context) error { return echo.ErrNotFound }

// Patch 修改一个对象的属性
func (api *ProductCategRelationAPI) Patch(c echo.Context) error { return echo.ErrNotFound }

// Delete 删除一个对象
func (api *ProductCategRelationAPI) Delete(c echo.Context) error {
	rel := &model.ProductCategRelation{UUID: c.Param("uuid")}
	ok, err := rel.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
