package routes

import "go-course/api/controllers"

var usersRoutes = []Route{
	{
		Method:       "POST",
		Path:         "/signup",
		RouteHandler: controllers.Signup,
	},
	// {
	// 	Method:       "GET",
	// 	Path:         "/users",
	// 	RouteHandler: controllers.GetUsers,
	// },
	// {
	// 	Method:       "GET",
	// 	Path:         "/users/:id",
	// 	RouteHandler: controllers.GetUser,
	// },
	// {
	// 	Method:       "POST",
	// 	Path:         "/users",
	// 	RouteHandler: controllers.CreateUser,
	// },
	// {
	// 	Method:       "PUT",
	// 	Path:         "/users/:id",
	// 	RouteHandler: controllers.UpdateUser,
	// },
	// {
	// 	Method:       "DELETE",
	// 	Path:         "/users/:id",
	// 	RouteHandler: controllers.DeleteUser,
	// },
}
