package main

import (
	"github.com/sirupsen/logrus"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
	"github.com/followgo/myadmin/util/mylogrus"
)

func main() {
	// 加载配置文件
	if err := LoadConfig(); err != nil {
		logrus.WithError(err).Fatalln("加载配置文件")
	}
	logrus.Infoln("已经加载配置文件")

	// 初始化日志系统
	initLogger()

	// 初始化 ORM
	if err := orm.InitOrmAndSyncModels(
		new(model.Admin), new(model.File),
		new(model.Setting), new(model.Part), new(model.Banner), new(model.Partner),
		new(model.ArticleCategory), new(model.Article),
		new(model.MarketSegment), new(model.TypicalCase), new(model.Solution),
		new(model.ProductCategory), new(model.Product), new(model.ProductImageRelation), new(model.ProductDetailPart), new(model.ProductCategoryRelation),
	); err != nil {
		logrus.WithError(err).Fatalln("初始化ORM并同步数据模型")
	}
	logrus.Infoln("已经初始化ORM和同步数据模型")

	// 尝试插入初始化数据
	tryInsertFactoryDefaultData()

	// 初始化并启动HTTP服务
	startHTTPServer()
}

// tryInsertFactoryDefaultData 尝试插入初始化数据
func tryInsertFactoryDefaultData() {
	// 插入默认的超级用户
	var u = new(model.Admin)
	if n, err := u.Count(nil); err != nil {
		logrus.WithError(err).Fatalln("获取用户数量")
	} else if n == 0 {
		u = &model.Admin{Username: "admin", Email: "admin@local", Password: "admin", Roles: []string{"viewer", "editor", "admin"}, Enabled: true}
		if ok, err := u.Insert(); err != nil {
			logrus.WithError(err).Fatalln("插入默认的超级用户")
		} else if !ok {
			logrus.Fatalln("失败插入默认的超级用户")
		}
	}

}

// initLogger 初始化日志系统
func initLogger() {
	logOpt := mylogrus.DefaultOption
	logOpt.BaseFile = Cfg.Logger.File
	logOpt.OverWrite = Cfg.Logger.OverWrite
	logOpt.OutputConsole = false

	var err error
	logOpt.Level, err = logrus.ParseLevel(Cfg.Logger.Level)
	if err != nil {
		logrus.WithError(err).Warnln("解析日志等级")
		logOpt.Level = logrus.InfoLevel
	}

	mylogrus.SetStdLogrus(logOpt)
	logrus.Infoln("已经初始化主日志记录器")
}
