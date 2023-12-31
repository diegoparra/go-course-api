package routes

import "go-course/api/controllers"

var eventsRoutes = []Route{
	{
		Method:       "GET",
		Path:         "/events",
		RouteHandler: controllers.GetEvents,
	},
	{
		Method:       "GET",
		Path:         "/events/:id",
		RouteHandler: controllers.GetEvent,
	},
	{
		Method:       "POST",
		Path:         "/events",
		RouteHandler: controllers.CreateEvent,
	},
	{
		Method:       "PUT",
		Path:         "/events/:id",
		RouteHandler: controllers.UpdateEvent,
	},
	{
		Method:       "DELETE",
		Path:         "/events/:id",
		RouteHandler: controllers.DeleteEvent,
	},
}
