package routers

import (
	"{{package}}/controllers"
)

func getUserRoutes() (routes []Route) {
	//type HandlerFunc func(*Context)
	routes = []Route{
		{"Index", "GET", "/users", controllers.UserList},
		{"get user", "GET", "/user/:id", controllers.UserIndex},
		{"SaveUserHandle", "POST", "/user", controllers.SaveUser},
	}

	return routes
}
