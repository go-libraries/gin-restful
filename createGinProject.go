package main

import (
	"flag"
	"gin-restful/gen"
)

var (
	port        string
	dsn         string
	projectName string
	projectPath string
	help        bool
	h           bool
)

func init() {
	flag.StringVar(&dsn, "dsn", "", "connection info names dsn")
	flag.StringVar(&projectPath, "path", "", "project build in this path")
	flag.StringVar(&projectName, "package", "", "package name use all project")
	flag.StringVar(&port, "port", "", "port")
	flag.BoolVar(&help, "help", false, "this help")
	flag.BoolVar(&h, "h", false, "this help")
}

func main() {

	flag.Parse()
	if h {
		flag.Usage()
	}
	if port == "" {
		port = "8080"
	}
	if projectName == "" || projectPath == "" {
		flag.Usage()
		return
	}
	flag.Usage = usage
	build := &gen.Build{
		TemplateFiles: make(map[string]string),
	}
	build.ProjectPath =  projectPath//"/Users/limars/Desktop/mango/ttt"
	build.ProjectName = projectName
	build.Load("")
	build.Generator()

	gen.BuildModels(build)
}

func usage() {
	flag.PrintDefaults()
}
