package config

// JWT 相关
const (
	// TokenSigningKey 用于 token 签名
	TokenSigningKey = "g9RpCZ8c1Cte7Y5iN=UBEaGz"

	// TokenAuthScheme Token字符串的前缀
	TokenAuthScheme = "vCD2cx"

	// TokenContextKey 将 Token 转存到 echo 的上下文对象
	TokenContextKey = "My-JWT-Token-X"

	// TokenLookup token请求的携带方式
	TokenLookup = "header:JWT-Token-X"
)
