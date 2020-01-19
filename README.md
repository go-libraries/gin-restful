# gin http restful项目生成

# 快速入门 
    
    go get  github.com/go-libraries/gin-restful

    //export GO111MODULE=on && export GOMOD={{pwd}}\go.mod && go install createGinProject.go
    
    $GOBIN/gin-restful  -package=项目名(包名) -path=项目路径 -dsn="username:password@tcp(host:port)/database"
    
# 生成项目基本结构

```go
-config
  config.ini
-controllers
  http处理回调
-logs
  日志文件目录
-models
  模型文件
-services
  主要对控制器的复杂业务二次封装
-routers
  路由文件
go.mod
main.go
Readme.md
```

# 参数详解

```go
Usage of createMangoProject:
  -dsn string
        connection info names dsn
  -h    this help
  -help
        this help
  -package string
        package name use all project
  -path string
        project build in this path
  -port string
        port
```

# 二次开发详解

## 控制器

1. 可以在services中书写新的控制器

```go
type UserSaveService struct {
	base.Controller
	Account *UserSaveRequest
}
func (u *UserSaveService) Decode() base.IError {
    // 解析 输入字段 如下
	u.Account = &UserSaveRequest{&models.UserAccount{}, ""}
	if bt, err := u.Ctx.GetRawData(); err == nil {
		if err := json.Unmarshal(bt, u.Account); err != nil {
			return base.NewError(err)
		}
	} else {
		return base.NewError(err)
	}

	return nil
}

func (u *UserSaveService) Process() base.IError {
    //todo:执行业务过程
	return nil
}
```

2. 可以在controllers中注入执行方法

```go
func SaveUser(c *gin.Context) {
	p := &base.Controller{}
	p.ServiceFun = func(u *base.DefaultService) base.IError {
		u.Data = "hello world"
		return nil
	}
	base.RunService(p, c)
}
```

## 路由

路由可以开发二次中间件功能
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"{{package}}/base"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func calTime(fn func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		fn(c)
		base.Log.Printf("Done in %v (%s %s)\n", time.Since(start), c.Request.Method, c.Request.URL.Path)
	}
}

func init() {

	//Router.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})

	for _, route := range getUserRoutes() {
		handle := calTime(route.HandlerFunc)
		base.Gin.Handle(route.Method, route.Path, handle)
	}

	//todo: add other Routes
}
```

## 模型

默认使用gorm作为数据驱动，如果初始化--dsn项目不为空，会自动将该db下表生成模型并提供基础方法

外部库详见 [外部库-模型生成器](https://github.com/go-libraries/genModels)


## 文档

文档使用swagger进行配置，可以一键生成

详见