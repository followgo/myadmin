package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/apiv1"
)

// RegisterAPIv1 注册 API 到指定的 URL
func RegisterAPIv1(e *echo.Echo) {
	g := e.Group("/api/v1")

	// 测试
	g.GET("/helloworld", apiv1.HelloWorld)

	// 不需要登陆
	g.POST("/login", new(apiv1.LoginAPI).LoginByLocal)
	g.POST("/ldap/login", new(apiv1.LoginAPI).LoginByLDAP)
	g.GET("/images/:uuid", new(apiv1.FileAPI).GetImage)

	// 需要 token 才能访问
	// mw.UsePermission(g)

	// 其它
	g.POST("/logout", new(apiv1.LoginAPI).Logout)
	g.POST("/refresh_token", new(apiv1.LoginAPI).RefreshToken)
	registerAPI(g, "/admins", new(apiv1.AdminAPI), nil, nil)
	registerAPI(g, "/files", new(apiv1.FileAPI), nil, nil)
	registerAPI(g, "/company/news", new(apiv1.CompanyNewsAPI), nil, nil)
	registerAPI(g, "/banners", new(apiv1.BannerAPI), nil, nil)

	// 部件和设置
	registerAPI(g, "/settings", new(apiv1.SettingAPI), nil, nil)
	registerAPI(g, "/parts", new(apiv1.PartAPI), nil, nil)

	// 文章
	registerAPI(g, "/article/categories", new(apiv1.ArticleCategAPI), nil, nil)
	registerAPI(g, "/articles", new(apiv1.ArticleAPI), nil, nil)

	// 解决方案和典型案例
	registerAPI(g, "/market/segments", new(apiv1.MarketSegmentAPI), nil, nil)
	registerAPI(g, "/solutions", new(apiv1.SolutionAPI), nil, nil)
	registerAPI(g, "/typical_cases", new(apiv1.TypicalCaseAPI), nil, nil)

	// 产品
	registerAPI(g, "/product/categories", new(apiv1.ProductCategAPI), nil, nil)
	registerAPI(g, "/product/categ/relations", new(apiv1.ProductImageRelationAPI), nil, nil)
	registerAPI(g, "/product/image/relations", new(apiv1.ProductImageRelationAPI), nil, nil)
	registerAPI(g, "/product/detail_parts", new(apiv1.ProductDetailPartAPI), nil, nil)
	registerAPI(g, "/products", new(apiv1.ProductAPI), nil, nil)
}
