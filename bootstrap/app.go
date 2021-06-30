package bootstrap

import (
	"goto/app/models/user"
	"goto/config"
	"goto/database/mysql"
	"goto/database/redis"
	"goto/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig("config")
	mysql.Conn()
	redis.Conn()
}

func Start() *gin.Engine {
	user.Init()
	return routes.Route
}
