package main

import (
	"os"

	"github.com/sirupsen/logrus"

	. "github.com/followgo/myadmin/config"
)

// init 创建初始化对象
func init() {
	if err := os.MkdirAll(Cfg.Upload.Directory, 0755); err != nil {
		logrus.WithError(err).Fatalln("创建文件上传目录")
	}
}
