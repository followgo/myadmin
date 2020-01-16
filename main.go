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
	if err := orm.InitOrm(); err != nil {
		logrus.WithError(err).Fatalln("初始化ORM")
	}
	logrus.Infoln("已经初始化ORM")

	// 同步数据模型结构体到数据库表
	if err := orm.Orm.Sync2(
		new(model.User),
	); err != nil {
		logrus.WithError(err).Fatalln("同步数据模型")
	}
	logrus.Infoln("已经同步数据模型")

	// 初始化并启动HTTP服务
	startHTTPServer()
}

// initLogger 初始化日志系统
func initLogger() {
	logOpt := mylogrus.DefaultOption
	logOpt.BaseFile = C.Logger.File
	logOpt.OverWrite = C.Logger.OverWrite
	logOpt.OutputConsole = false

	var err error
	logOpt.Level, err = logrus.ParseLevel(C.Logger.Level)
	if err != nil {
		logrus.WithError(err).Warnln("解析日志等级")
		logOpt.Level = logrus.InfoLevel
	}

	mylogrus.SetStdLogrus(logOpt)
	logrus.Infoln("已经初始化主日志记录器")
}
