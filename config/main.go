package config

import (
	"os"
	"path/filepath"

	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/configurator"
)

const (
	// configFile 配置文件路径
	configFile = "./config/config.toml"
)

// LoadConfig 装载配置文件
func LoadConfig() error {
	c := configurator.NewConfigurator(configFile, Cfg)

	if has, err := util.HasFile(configFile); err != nil {
		return err
	} else if !has {
		if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
			return err
		}

		return c.Save("配置文件，请谨慎修改")
	}

	return c.Load()
}
