package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("读取 %s 配置文件:%s 失败", configName, err)
	}
}
