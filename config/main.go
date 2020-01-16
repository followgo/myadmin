package config

import (
	"os"
	"path/filepath"

	"github.com/followgo/myadmin/util/configurator"
	"github.com/followgo/myadmin/util/file"
)

const (
	// configFile 配置文件路径
	configFile = "./config/config.toml"
)

// LoadConfig 装载配置文件
func LoadConfig() error {
	c := configurator.NewConfigurator(configFile, C)

	if has, err := file.HasFile(configFile); err != nil {
		return err
	} else if !has {
		if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
			return err
		}

		return c.Save("配置文件")
	}

	return c.Load()
}
