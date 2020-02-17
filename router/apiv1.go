package router

import (
	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/router/apiv1"
)

func RegisterAPIv1(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// 测试
	v1.GET("/helloworld", apiv1.HelloWorld)

	// 不需要登陆
	v1.POST("/login", new(apiv1.LoginAPI).LoginByLocal)
	v1.POST("/ldap/login", new(apiv1.LoginAPI).LoginByLDAP)
	v1.GET("/images/:uuid", new(apiv1.FileAPI).GetImage)

	// 需要 token 才能访问
	// mw.UsePermission(v1)

	// 其它
	v1.POST("/logout", new(apiv1.LoginAPI).Logout)
	v1.POST("/refresh_token", new(apiv1.LoginAPI).RefreshToken)
	registerAPI(v1, "/admins", new(apiv1.AdminAPI), nil, nil)
	registerAPI(v1, "/files", new(apiv1.FileAPI), nil, nil)
	registerAPI(v1, "/company/news", new(apiv1.CompanyNewsAPI), nil, nil)
	registerAPI(v1, "/banners", new(apiv1.BannerAPI), nil, nil)

	// 部件和设置
	registerAPI(v1, "/settings", new(apiv1.SettingAPI), nil, nil)
	registerAPI(v1, "/parts", new(apiv1.PartAPI), nil, nil)

	// 文章
	registerAPI(v1, "/article/categories", new(apiv1.ArticleCategAPI), nil, nil)
	registerAPI(v1, "/articles", new(apiv1.ArticleAPI), nil, nil)

	// 解决方案和典型案例
	registerAPI(v1, "/market/segments", new(apiv1.MarketSegmentAPI), nil, nil)
	registerAPI(v1, "/solutions", new(apiv1.SolutionAPI), nil, nil)
	registerAPI(v1, "/typical_cases", new(apiv1.TypicalCaseAPI), nil, nil)

	// 产品
	registerAPI(v1, "/product/categories", new(apiv1.ProductCategAPI), nil, nil)
	registerAPI(v1, "/product/categ/relations", new(apiv1.ProductImageRelationAPI), nil, nil)
	registerAPI(v1, "/product/image/relations", new(apiv1.ProductImageRelationAPI), nil, nil)
	registerAPI(v1, "/product/detail_parts", new(apiv1.ProductDetailPartAPI), nil, nil)
	registerAPI(v1, "/products", new(apiv1.ProductAPI), nil, nil)
}
