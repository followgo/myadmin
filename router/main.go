package router

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// API 定义结构化的 API
type API interface {
	// Get 根据 `uuid` 获取某个指定的对象，向客户端返回单个资源对象
	// HTTP Method: GET, SuccessStatusCode=200
	Get(c echo.Context) error

	// Select 列出所有对象，支持 Filtering，向客户端返回资源对象的列表
	// HTTP Method: GET, SuccessStatusCode=200, 202
	Select(c echo.Context) error

	// Create 新建一个对象，向客户端返回完整的资源对象
	// HTTP Method: POST, SuccessStatusCode=201
	Create(c echo.Context) error

	// Update 更新对象（客户端提供改变后的完整资源），向客户端返回完整的资源对象
	// HTTP Method: PUT, SuccessStatusCode=201
	Update(c echo.Context) error

	// Patch 更新对象的特定属性（客户端提供改变的属性），向客户端返回完整的资源对象
	// HTTP Method: PATCH, SuccessStatusCode=201
	Patch(c echo.Context) error

	// Delete 删除一个对象，向客户端返回一个空文档
	// HTTP Method: DELETE, SuccessStatusCode=204
	Delete(c echo.Context) error
}

// registerAPI 注册 API，根据 RESTful 风格注册路径
func registerAPI(g *echo.Group, urlPath string, api API, readMWFuncs, writeMWFuncs []echo.MiddlewareFunc) {
	urlPath = strings.TrimRight(urlPath, "/")

	if readMWFuncs == nil {
		g.GET(urlPath+"/:uuid", api.Get)
		g.GET(urlPath, api.Select)
	} else {
		g.GET(urlPath+"/:uuid", api.Get, readMWFuncs...)
		g.GET(urlPath, api.Select, readMWFuncs...)
	}

	if writeMWFuncs == nil {
		g.POST(urlPath, api.Create)
		g.PUT(urlPath+"/:uuid", api.Update)
		g.PATCH(urlPath+"/:uuid", api.Patch)
		g.DELETE(urlPath+"/:uuid", api.Delete)
	} else {
		g.POST(urlPath, api.Create, writeMWFuncs...)
		g.PUT(urlPath+"/:uuid", api.Update, writeMWFuncs...)
		g.PATCH(urlPath+"/:uuid", api.Patch, writeMWFuncs...)
		g.DELETE(urlPath+"/:uuid", api.Delete, writeMWFuncs...)
	}
}
