# 芒果听见gin http项目生成

目前无任何外部库，直接可用

# 快速入门 

    git clone https://github.com/go-libraries/gin-restful
    
    cd gin-restful
	
    go mod download
    
    go install createGinProject.go
    
    $GOBIN/createGinProject  -package=项目名 -path=项目路径
    
# 生成项目基本结构

```go
-base
  基础文件
-config
  config.ini
-handlers
  http处理回调
-logs
  日志文件目录
-models
  模型文件
-processes
  可以理解为控制器对象
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
