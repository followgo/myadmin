package config

import (
	"github.com/followgo/myadmin/util/imagex"
	"github.com/followgo/myadmin/util/random"
)

// Cfg 存储配置
var Cfg = &config{
	SecuritySalt: random.String(24),

	HTTP: httpConfig{
		Debug:        true,
		ListenAddr:   "127.0.0.1:1213",
		AccessFile:   "./log/access.log",
		AllowOrigins: []string{"*"},
	},
	Logger: loggerConfig{
		Level:     "INFO",
		File:      "./log/main.log",
		OverWrite: false,
	},
	Orm: ormConfig{
		Debug:       true,
		DriverName:  "sqlite3",
		DriverUri:   "./main.s3db",
		LogLevel:    "INFO",
		LogFile:     "./log/orm.log",
		UseLRUCache: false,
	},
	Upload: uploadConfig{
		Directory: "./upload",
		AllowMIMETypes: []string{
			imagex.MIMETypeByExtension(".webp"),
			imagex.MIMETypeByExtension(".png"),
			imagex.MIMETypeByExtension(".jpg"),
			imagex.MIMETypeByExtension(".gif"),
			imagex.MIMETypeByExtension(".pdf"),
			imagex.MIMETypeByExtension(".zip"),
			imagex.MIMETypeByExtension(".webm"),
		},
		AllowMaxSizeMB:       10,
		ConvertPictureToWebp: true,
	},
	LDAP: ldapConfig{
		Enabled:                false,
		Protocol:               "tcp",
		UseTLS:                 true,
		StartTLS:               true,
		ServerAddr:             "ldap.example.org:636",
		BindSearcherDN:         "cn=Searcher,ou=IT,dc=example,dc=org",
		BindSearcherDNPassword: "123456",
		SearchBaseDN:           "ou=IT,dc=example,dc=org",
		SearchFilter:           "(&(objectClass=organizationalPerson))",
		RDNAttr:                "cn",
		UserAttributes: ldapUserAttributes{
			CommonName:  "cn",
			Surname:     "sn",
			Telephone:   "telephoneNumber",
			Description: "description",
		},
	},
}

// config 主配置
type config struct {
	// SecuritySalt 安全盐，用于密码哈希
	SecuritySalt string

	// HTTP 配置
	HTTP httpConfig

	// Logger 日志配置
	Logger loggerConfig

	// Orm ORM 引擎配置
	Orm ormConfig

	// Upload 上传设置
	Upload uploadConfig

	// ldap LDAP认证的设置
	LDAP ldapConfig `json:"ldap"`
}

// httpConfig HTTP 配置
type httpConfig struct {
	// Debug 调试模式，禁用 Recover 中间件
	Debug bool

	// ListenAddr 监听地址
	ListenAddr string

	// AccessFile 会话记录文件
	AccessFile string

	// AllowOrigins 允许的跨域
	AllowOrigins []string
}

// loggerConfig 日志配置
type loggerConfig struct {
	// File 日志文件
	File string

	// Level 日志级别
	Level string

	// OverWrite 覆盖旧文件
	OverWrite bool
}

// ormConfig ORM 引擎配置
type ormConfig struct {
	// Debug 调试模式
	Debug bool

	// DriveName 数据库
	DriverName string

	// DriverUri 数据库连接
	DriverUri string

	// LogLevel 日志级别
	LogLevel string

	// LogFile 日志记录文件
	LogFile string

	// UseLRUCache 开启缓存
	UseLRUCache bool
}

// uploadConfig 上传设置
type uploadConfig struct {
	// Directory 存储目录
	Directory string

	// AllowMIMETypes 允许上传的文件类型
	AllowMIMETypes []string

	// AllowMaxSizeMB 允许上传的最大尺寸
	AllowMaxSizeMB int64

	// ConvertPictureToWebp 转换图片格式为webp格式
	ConvertPictureToWebp bool
}

// ldapConfig LDAP登陆设置
type ldapConfig struct {
	// Enabled 使能 LDAP 认证功能
	Enabled bool

	// Protocol 使用协议，TCP 或 UDP
	Protocol string

	// UseTLS 使用 LDAP with SSL/TLS
	UseTLS   bool
	StartTLS bool

	// 服务器地址
	ServerAddr string

	// SearchDN 绑定用于查询的用户
	BindSearcherDN         string
	BindSearcherDNPassword string

	// SearchFilter 查询范围和匹配过滤
	SearchBaseDN string
	SearchFilter string

	// RDNAttr RDN(Relative Distinguished Name)对应的属性名称，通常是 `cn` 或者 `uid`
	RDNAttr string
	// Attributes 用户的属性字段名称
	UserAttributes ldapUserAttributes
}

// ldapUserAttributes LDAP 用户的属性字段名称
type ldapUserAttributes struct {
	CommonName  string // 名，person对象的必须要属性 cn
	Surname     string // 姓，person对象的必须要属性 sn
	Telephone   string // 电话，person对象的可选属性 telephoneNumber
	Description string // 描述，person对象的可选属性 description
}
