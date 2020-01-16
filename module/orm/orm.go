package orm

import (
	"log"
	"strings"

	"github.com/xormplus/core"
	"github.com/xormplus/xorm"

	. "github.com/followgo/myadmin/config"

	"github.com/followgo/myadmin/util/errors"
	"github.com/followgo/myadmin/util/mylogrus"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Orm 引擎接口
	Orm xorm.EngineInterface
)

// InitOrm 初始化ORM
func InitOrm() (err error) {
	Orm, err = xorm.NewEngine(C.Orm.DriverName, C.Orm.DriverUri)
	if err != nil {
		return errors.Wrap(err, "创建 xorm 实例")
	}

	// 配置日志记录器
	logWriter := mylogrus.NewWriterWithSizeRotate(C.Orm.LogFile, 100, 100, 30)
	Orm.SetLogger(&xorm.SimpleLogger{
		DEBUG: log.New(logWriter, "[D] ", xorm.DEFAULT_LOG_FLAG),
		ERR:   log.New(logWriter, "[E] ", xorm.DEFAULT_LOG_FLAG),
		INFO:  log.New(logWriter, "[I] ", xorm.DEFAULT_LOG_FLAG),
		WARN:  log.New(logWriter, "[W] ", xorm.DEFAULT_LOG_FLAG),
	})
	Orm.ShowExecTime(C.Orm.Debug)
	Orm.ShowSQL(C.Orm.Debug)
	switch strings.ToLower(C.Orm.LogLevel) {
	case "warn", "warning":
		Orm.SetLogLevel(core.LOG_WARNING)
	case "err", "error":
		Orm.SetLogLevel(core.LOG_ERR)
	case "info":
		Orm.SetLogLevel(core.LOG_INFO)
	default:
		Orm.SetLogLevel(core.LOG_DEBUG)
	}

	if C.Orm.UseLRUCache {
		enableLRUCacher()
	}

	return nil
}

// enableLRUCacher 启用缓存
// 当使用了Distinct,Having,GroupBy方法将不会使用缓存
// 避免使用 Exec 执行写操作，如果执行了要清理缓存，如：engine.ClearCache(new(User))
func enableLRUCacher() {
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Orm.SetDefaultCacher(cacher)
}
