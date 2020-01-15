package routers

import "{{package}}/controllers"

func GroupExample()  *GroupRoutes  {
	g := getNewGroupRoutes("/group","/group")
	
	g.addRoute(Route{
		Name:        "list",
		Method:      "GET",
		Path:        "/list",
		HandlerFunc: controllers.UserList,
	})

	return g
}