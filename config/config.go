package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(envPath string) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("读取 .env 配置文件:%s 失败", err)
	}else{
		logrus.Info("load config success")
	}
}
