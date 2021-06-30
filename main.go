package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"goto/bootstrap"
)

func main() {
	app := bootstrap.Start()
	addr := fmt.Sprintf(":%s", viper.GetString("app.port"))
	if err := app.Run(addr); err != nil {
		logrus.Error(err)
	}
}
