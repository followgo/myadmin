package config

import (
	"github.com/followgo/myadmin/util/random"
)

const (
	// TokenContextKey 将 token 转存到 echo 的上下文对象
	TokenContextKey = "x-token"
)

var (
	// TokenSigningKey 用于 token 签名
	TokenSigningKey = random.String(24)
)