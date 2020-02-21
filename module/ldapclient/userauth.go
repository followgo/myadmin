package ldapclient

import (
	"fmt"
	"strings"

	. "github.com/followgo/myadmin/config"
)

// UserAuth 查询用户信息，验证用户的密码
func UserAuth(dn, password string) (map[string]string, error) {
	// 连接服务器
	conn, err := dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 查询用户信息
	dnSlice := strings.Split(dn, ",")
	filter := Cfg.LDAP.SearchFilter[:len(Cfg.LDAP.SearchFilter)-1] + fmt.Sprintf("(%s)", dnSlice[0]) + ")"
	users, err := search(conn, strings.Join(dnSlice[1:], ","), filter)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrNoFound
	}

	userInfos := make(map[string]string)
	for _, u := range users {
		if strings.ToLower(u["dn"]) == strings.ToLower(dn) {
			userInfos = u
		}
	}
	if len(userInfos) == 0 {
		return nil, ErrNoFound
	}

	// 验证用户的密码
	if err := conn.Bind(userInfos["dn"], password); err != nil {
		return nil, ErrUserAuth
	}

	return userInfos, nil
}
