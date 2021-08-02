package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	// 初始化配置文件
	viper.SetConfigFile("config/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
