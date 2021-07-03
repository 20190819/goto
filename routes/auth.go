package routes

import "goto/app/http/controllers"

func RouterAuth() {
	// 路由注册
	routeItems := make([]Item, 2)
	authController := controllers.AuthController{}

	routeItems = append(routeItems, Item{Method: "post", Url: "register", Action: authController.Register})
	routeItems = append(routeItems, Item{Method: "post", Url: "login", Action: authController.Login})

	Register(routeItems, RouteGroup)
}
