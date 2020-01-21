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
