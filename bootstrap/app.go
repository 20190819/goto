package bootstrap

import (
	"goto/app/models/user"
	"goto/config"
	"goto/database/mysql"
	"goto/database/redis"
	"goto/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig("./")
}

func Start() *gin.Engine {
	mysql.Conn()
	redis.Conn()
	routes.Route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello my goto framework")
	})
	user.Migration()
	routes.RouterAuth()
	return routes.Route
}
