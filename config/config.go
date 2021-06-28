package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigFile("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("读取 %s 配置文件:%s 失败", configName, err)
		os.Exit(1)
	}
}
