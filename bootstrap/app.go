package bootstrap

import (
	"goto/app/models/user"
	"goto/config"
	"goto/database/mysql"
	"goto/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig(".env")
}

func Start() *gin.Engine {
	mysql.Conn()
	//redis.Conn()
	routes.Route.GET("/home", func(context *gin.Context) {
		context.String(http.StatusOK, "hello my goto framework")
	})
	user.Init()
	return routes.Route
}
