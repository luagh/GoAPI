package main

import (
	"GOHUB/app/cmd"
	"GOHUB/bootstrap"
	btsConfig "GOHUB/config"
	"GOHUB/pkg/config"
	"GOHUB/pkg/console"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
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
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
