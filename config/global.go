package config

import (
	"github.com/followgo/myadmin/util/random"
)

// JWT 相关
var (
	// TokenSigningKey 用于 token 签名
	TokenSigningKey = random.String(24)

	// TokenAuthScheme Token字符串的前缀
	TokenAuthScheme = random.String(6)

	// TokenContextKey 将 Token 转存到 echo 的上下文对象
	TokenContextKey = "My-JWT-Token-X"

	// TokenLookup token请求的携带方式
	TokenLookup = "header:JWT-Token-X"
)
