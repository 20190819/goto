package bootstrap

import (
	"goto/config"

	"goto/database/mysql"
	"goto/database/redis"
)

func init() {
	config.InitConfig("config")
	mysql.Conn()
	redis.Conn()
}

func Start() {
	
}
