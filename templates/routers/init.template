package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grestful/core"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

type GroupRoutes struct {
	Name	string
	Path    string
	Routes  []Route
}

func getNewGroupRoutes(name, path string) *GroupRoutes {
	return &GroupRoutes{
		Name:   name,
		Path:   path,
		Routes: make([]Route,0),
	}
}

func (gr *GroupRoutes) addRoute(r Route) {
	gr.Routes = append(gr.Routes, r)
}



func calTime(fn func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		fn(c)
		core.GetLog().Printf("Done in %v (%s %s)\n", time.Since(start), c.Request.Method, c.Request.URL.Path)
	}
}

func middlewareExample(c *gin.Context)  {
	fmt.Println("我是goods路由组中间件1")
}

func init() {

	//Router.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})
	base := core.GetCore()
	for _, route := range getUserRoutes() {
		handle := calTime(route.HandlerFunc)
		base.Gin.Handle(route.Method, route.Path, handle)
	}

	g := GroupExample()
	gr := base.Gin.Group(g.Path)

	gr.Use(middlewareExample)
	for _, route := range g.Routes {
		gr.Handle(route.Method, route.Path, route.HandlerFunc)
	}
	//todo: add other Routes
}
