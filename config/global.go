package config

import (
	"github.com/followgo/myadmin/util/random"
)

const (
	// TokenContextKey 将 Token 转存到 echo 的上下文对象
	TokenContextKey = "X-Token"

	// TokenAuthScheme Token字符串的前缀
	TokenAuthScheme = "eyJhbGciOi"

	// TokenLookup token请求的携带方式
	TokenLookup = "header:JWT-Token-X"
)

var (
	// TokenSigningKey 用于 token 签名
	TokenSigningKey = random.String(24)
)
