package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// ProductImageRelationAPI 产品图片API
type ProductImageRelationAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *ProductImageRelationAPI) Get(c echo.Context) error {
	pImg := &model.ProductImageRelation{UUID: c.Param("uuid")}
	has, err := pImg.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, pImg)
}

// Select 列出所有选择的对象
func (api *ProductImageRelationAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	pImages, err := new(model.ProductImageRelation).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.ProductImageRelation).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "images": pImages})
}

// Create 创建一个新对象
func (api *ProductImageRelationAPI) Create(c echo.Context) error {
	pImg := &model.ProductImageRelation{}
	if err := c.Bind(pImg); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := pImg.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, pImg)
}

// Update 完全更新一个对象
func (api *ProductImageRelationAPI) Update(c echo.Context) error { return echo.ErrNotFound }

// Patch 修改一个对象的属性
func (api *ProductImageRelationAPI) Patch(c echo.Context) error { return echo.ErrNotFound }

// Delete 删除一个对象
func (api *ProductImageRelationAPI) Delete(c echo.Context) error {
	pImg := &model.ProductImageRelation{UUID: c.Param("uuid")}
	ok, err := pImg.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
