package bootstrap

import "goto/config"

func init() {

}

func Start() {
	config.InitConfig("config")
	mysqlInit()
}
