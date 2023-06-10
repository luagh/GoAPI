package cmd

import (
	"GOHUB/bootstrap"
	"GOHUB/pkg/config"
	"GOHUB/pkg/console"
	"GOHUB/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 将程序修改为命令模式
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例 处理 HTTP 请求和路由
	router := gin.New()

	// 初始化路由绑定将不同的 URL 映射到相应的处理函数上
	bootstrap.SetupRoute(router)

	// 运行 启动 Web服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
