package user

import (
	"goto/app/http/controllers"
	"goto/database/mysql"
	"goto/routes"
)

func Init() {
	// 迁移表结构
	migrates := make(map[string]interface{})
	migrates["user"] = User{}
	mysql.AutoMigrate(migrates)
	// 路由注册
	routeItems := make([]routes.Item, 2)
	authController := controllers.AuthController{}

	routeItems = append(routeItems, routes.Item{Method: "post", Url: "register", Action: authController.Register})
	routeItems = append(routeItems, routes.Item{Method: "post", Url: "login", Action: authController.Login})

	routes.Register(routeItems, routes.AuthRouteGroup)

}
