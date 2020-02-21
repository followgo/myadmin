package ldapclient

import (
	"github.com/followgo/myadmin/util/errors"
)

var (
	// ErrNoFound 没有找到任何用户
	ErrNoFound = errors.New("no users found")

	// ErrUserAuth 用户认证失败
	ErrUserAuth = errors.New("user authentication failed")
)
