package main

import (
	"os"

	"github.com/sirupsen/logrus"

	. "github.com/followgo/myadmin/config"
)

// init 创建初始化对象
func init() {
	for _, d := range []string{"./log", Cfg.Upload.Directory} {
		if err := os.MkdirAll(d, 0755); err != nil {
			logrus.Fatalln(err)
		}
	}
}
