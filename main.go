package main

import (
	"GOHUB/bootstrap"
	btsConfig "GOHUB/config"
	"GOHUB/pkg/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 初始化 Redis
	bootstrap.SetupRedis()
	//  // 初始化路由绑定
	bootstrap.SetupRoute(r)
	// 初始化 Logger
	bootstrap.SetupLogger()
	gin.SetMode(gin.ReleaseMode)
	//初始化DB
	bootstrap.SetupDB()
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}

}
