package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
)

// ArticleCategAPI 文章类别API
type ArticleCategAPI struct{}

// Get 根据 `uuid` 获取指定的对象
func (api *ArticleCategAPI) Get(c echo.Context) error {
	categ := &model.ArticleCateg{UUID: c.Param("uuid")}
	has, err := categ.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}
	if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "没有此数据"}
	}

	return c.JSON(http.StatusOK, categ)
}

// Select 列出所有选择的对象
func (api *ArticleCategAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	categories, err := new(model.ArticleCateg).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	// 数量
	total, err := new(model.ArticleCateg).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "categories": categories})
}

// Create 创建一个新对象
func (api *ArticleCategAPI) Create(c echo.Context) error {
	categ := &model.ArticleCateg{}
	if err := c.Bind(categ); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	ok, err := categ.Insert()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, categ)
}

// Update 完全更新一个对象
func (api *ArticleCategAPI) Update(c echo.Context) error {
	categ := &model.ArticleCateg{}
	if err := c.Bind(categ); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	categ.UUID = c.Param("uuid")

	n, err := categ.Update(nil, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, categ)
}

// Patch 修改一个对象的属性
func (api *ArticleCategAPI) Patch(c echo.Context) error {
	categ := &model.ArticleCateg{}
	if err := c.Bind(categ); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	categ.UUID = c.Param("uuid")

	cols := strings.Split(c.QueryParam("cols"), ",")
	if len(cols) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "缺少必要的参数"}
	}

	n, err := categ.Update(cols, nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据出错", Internal: err}
	}
	if n != 1 {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "更新数据失败"}
	}

	return c.JSON(http.StatusCreated, categ)
}

// Delete 删除一个对象
func (api *ArticleCategAPI) Delete(c echo.Context) error {
	categ := &model.ArticleCateg{UUID: c.Param("uuid")}
	ok, err := categ.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}
	return c.NoContent(http.StatusNoContent)
}
