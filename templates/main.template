package main

import (
	"fmt"
	ginRest "github.com/grestful/core"
	"github.com/grestful/utils"


	//swaggerFiles "github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	_ "{{package}}/controllers"
	//_ "{{package}}/docs"
	_ "{{package}}/routers"

	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	path := utils.GetCurrentPath() + "/config/config.ini"
	fmt.Println(path)
	ginRest.InitConfig(path)
	core := ginRest.GetCore()
	port, _ := core.Config.GetValue("", "PORT")
	if port == "" {
		port = "8081"
	}


	env, _ := core.Config.GetValue("", "ENV")
	fmt.Println(env)
	//if env == "dev" {
	//	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)) // The url pointing to API definition
	//	core.Handle("GET","/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//}
	go func() {
		// 服务连接
		if err := core.Gin.Run(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
