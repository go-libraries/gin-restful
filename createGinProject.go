package main

import (
	"flag"
	"fmt"
	"gin-restful/gen"
	"os"
	"runtime"
	"strings"
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
	// dsn username:password@protocol(address)/dbname?param=value
	if projectName == "" || projectPath == "" || dsn == "" {
		flag.Usage()
		return
	}
	flag.Usage = usage
	build := &gen.Build{
		TemplateFiles: make(map[string]string),
	}
	build.ProjectPath = projectPath //"/Users/limars/Desktop/mango/ttt"
	build.ProjectName = projectName
	build.Dsn = dsn

	templatePath := ""
	goPath := os.Getenv("GOPATH")
	if goPath != "" {
		if runtime.GOOS == "windows" {
			goPaths := strings.Split(goPath, ";")
			for _,path := range goPaths  {
				_,err := os.Stat(path+"\\src\\github.com\\go-libraries\\gin-restful")
				if err == nil {
					templatePath = path+"\\src\\github.com\\go-libraries\\gin-restful"
					break
				}
			}
		}else{
			goPaths := strings.Split(goPath, ":")
			for _,path := range goPaths  {
				_,err := os.Stat(path+"/src/github.com/go-libraries/gin-restful")
				if err == nil {
					templatePath = path+"/src/github.com/go-libraries/gin-restful"
					break
				}
			}
		}
	}
	fmt.Println(templatePath)
	build.Load(templatePath)
	build.Generator()

	gen.BuildModels(build)
}

func usage() {
	flag.PrintDefaults()
}
