package apiv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/jwt"
	"github.com/followgo/myadmin/module/onlineuser"
)

type Login struct {
	// from1: 使用本地用户名和密码登陆系统
	from1 struct {
		Username string `json:"username" form:"username" query:"username"`
		Email    string `json:"email" form:"email" query:"email"`
		Password string `json:"password" form:"password" query:"password"`
	}

	// from2: 使用 LDAP 登陆系统
	from2 struct{}
}

// LoginByLocal 使用本地用户名和密码登陆登入系统
func (api *Login) LoginByLocal(c echo.Context) error {
	if err := c.Bind(api.from1); err != nil {
		return echo.ErrBadRequest
	}

	// 检验用户名和密码
	user := model.User{Username: api.from1.Username, Email: api.from1.Email, Password: api.from1.Password}
	if !user.Verify() {
		return echo.ErrUnauthorized
	}

	// 添加到在线用户列表
	if err := onlineuser.AddUser(onlineuser.User{
		UUID:      user.UUID,
		Username:  user.Username,
		Roles:     user.Roles,
		RemoteIP:  c.Request().RemoteAddr[:strings.LastIndexByte(c.Request().RemoteAddr, ':') ],
		UserAgent: c.Request().UserAgent(),
	}); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	// 创建 token 并发给用户
	token, err := jwt.GenerateTokenString(map[string]interface{}{"UUID": user.UUID})
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

// Logout 登出系统
func (api *Login) Logout(c echo.Context) error {
	claims := jwt.GetClaimsFromToken(c)
	onlineuser.RemoteUser(claims["UUID"].(string))
	return c.NoContent(http.StatusCreated)
}

// RefreshToken 刷新 token，更新 token 的签发时间，到期时间
func (api *Login) RefreshToken(c echo.Context) error {
	claims := jwt.GetClaimsFromToken(c)

	tokenStr, err := jwt.GenerateTokenString(claims)
	if err != nil {
		return echo.ErrInternalServerError
	}

	onlineuser.RefreshUser(claims["uuid"].(string))
	return c.JSON(http.StatusOK, echo.Map{"token": tokenStr})
}
