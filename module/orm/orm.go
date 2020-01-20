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
	// engine Orm 引擎接口
	engine xorm.EngineInterface
)

// NewSession 创建session，并设置声明
func NewSession(filter *Filter) (s *xorm.Session) {
	s = engine.NewSession()

	if filter == nil {
		return s
	}

	if filter.Cols != nil && len(filter.Cols) > 0 {
		s = s.Cols(filter.Cols...)
	} else {
		s = s.AllCols()
	}

	if filter.OmitCols != nil && len(filter.OmitCols) > 0 {
		s = s.Omit(filter.OmitCols...)
	}

	if filter.AscCols != nil && len(filter.AscCols) > 0 {
		s = s.Asc(filter.AscCols...)
	}

	if filter.DescCols != nil && len(filter.DescCols) > 0 {
		s = s.Desc(filter.DescCols...)
	}

	if filter.Query != "" {
		if filter.QueryArgs != nil && len(filter.QueryArgs) > 0 {
			s = s.Where(filter.Query, filter.QueryArgs...)
		} else {
			s = s.Where(filter.Query)
		}
	}

	if filter.GroupByKeys != "" {
		s = s.GroupBy(filter.GroupByKeys)
	}

	if filter.Limit[0] != 0 {
		s = s.Limit(filter.Limit[0], filter.Limit[1])
	}

	return s
}

// InitOrmAndSyncModels 初始化ORM，并同步数据模型
func InitOrmAndSyncModels(models ...interface{}) (err error) {
	engine, err = xorm.NewEngine(Cfg.Orm.DriverName, Cfg.Orm.DriverUri)
	if err != nil {
		return errors.Wrap(err, "创建 xorm 实例")
	}

	// 配置日志记录器
	logWriter := mylogrus.NewWriterWithSizeRotate(Cfg.Orm.LogFile, 100, 100, 30)
	engine.SetLogger(&xorm.SimpleLogger{
		DEBUG: log.New(logWriter, "[D] ", xorm.DEFAULT_LOG_FLAG),
		ERR:   log.New(logWriter, "[E] ", xorm.DEFAULT_LOG_FLAG),
		INFO:  log.New(logWriter, "[I] ", xorm.DEFAULT_LOG_FLAG),
		WARN:  log.New(logWriter, "[W] ", xorm.DEFAULT_LOG_FLAG),
	})
	engine.ShowExecTime(Cfg.Orm.Debug)
	engine.ShowSQL(Cfg.Orm.Debug)
	switch strings.ToLower(Cfg.Orm.LogLevel) {
	case "warn", "warning":
		engine.SetLogLevel(core.LOG_WARNING)
	case "err", "error":
		engine.SetLogLevel(core.LOG_ERR)
	case "info":
		engine.SetLogLevel(core.LOG_INFO)
	default:
		engine.SetLogLevel(core.LOG_DEBUG)
	}

	if Cfg.Orm.UseLRUCache {
		enableLRUCacher()
	}

	return engine.Sync2(models...)
}

// enableLRUCacher 启用缓存
// 当使用了Distinct,Having,GroupBy方法将不会使用缓存
// 避免使用 Exec 执行写操作，如果执行了要清理缓存，如：engine.ClearCache(new(User))
func enableLRUCacher() {
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)
}
