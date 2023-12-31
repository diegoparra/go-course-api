package routes

import "github.com/gin-gonic/gin"

type Route struct {
	Method       string
	Path         string
	RouteHandler func(*gin.Context)
}

func RegisterRoutes(server *gin.Engine) {
	routes := eventsRoutes
	routes = append(routes, usersRoutes...)

	for _, route := range routes {
		server.Handle(route.Method, route.Path, route.RouteHandler)
	}
}
