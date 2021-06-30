package routes

import (
	"github.com/gin-gonic/gin"
	"goto/app/http/middleware"
)

var Route = gin.Default()

var RouteGroup *gin.RouterGroup
var AuthRouteGroup *gin.RouterGroup

type Item struct {
	Method string
	Url    string
	Action gin.HandlerFunc
}

func init() {
	Route.Use(middleware.Cors())
	RouteGroup = Route.Group("api")
	AuthRouteGroup = Route.Group("api", middleware.ApiAuth())
}

func Group(relativePath string, middlewares ...gin.HandlerFunc) *gin.RouterGroup {
	routeGroup := Route.Group(relativePath)
	routeGroup.Use(middlewares...)
	return routeGroup
}

func Register(items []Item, group ...*gin.RouterGroup) {
	var route *gin.RouterGroup
	if len(group) > 0 {
		route = group[0]
	} else {
		route = RouteGroup
	}

	for _, item := range items {
		switch item.Method {
		case "get":
			route.GET(item.Url, item.Action)
		case "post":
			route.POST(item.Url, item.Action)
		case "put":
			route.PUT(item.Url, item.Action)
		case "delete":
			route.DELETE(item.Url, item.Action)
		}
	}
}
